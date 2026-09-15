// Package rfc3339 turns RFC 3339 text into nanoseconds since the Unix epoch.
//
// It is written in arithmetic alone, with no import but the api: the sandbox
// may not name `time`, and a timestamp may not cross the api as a time.Time
// either — an api type has to be convertible for a consumer to copy the
// contract into its own deps, and a type from another package is not. So an
// instant leaves this library the way it reaches any other agnos contract: as
// a plain int64 of nanoseconds, UTC.
package rfc3339

import (
	api "github.com/MateusMoutinhoOrg/Verb/sandbox/api"
)

// Layout is the timestamp spelling every Timestamp getter of api.Parser
// accepts, in the reference-time notation the Go standard library writes
// layouts with. It is here rather than in a config file because the grammar
// this package implements is RFC 3339 itself, not a format a project picks.
const Layout = "2006-01-02T15:04:05Z07:00"

// nanosPerSecond is the scale an instant is reported in.
const nanosPerSecond = 1000000000

// maxSeconds is the largest number of seconds either side of the epoch that
// still fits an int64 of nanoseconds once the sub-second part is added. It
// puts the representable range at roughly the years 1678 to 2262.
const maxSeconds = 9223372035

// ParseNano parses s as an RFC 3339 timestamp — "2024-01-02T15:04:05Z",
// "2024-01-02T15:04:05.123-03:00" — and returns it as nanoseconds since the
// Unix epoch, UTC. The offset is applied, so two spellings of one instant
// return the same number.
//
// It errors on anything the grammar does not allow: a field that is not all
// digits, a date that does not exist (31 April, 29 February of a common
// year), an hour, minute or second out of range, a missing or malformed zone,
// a zone offset outside ±23:59 — where the standard library is lenient and
// this is not — and an instant outside the range an int64 of nanoseconds can
// hold. A fractional second beyond the ninth digit is read and dropped, the
// way the standard library reads it.
func ParseNano(sandbox *api.Sandbox, text string) (int64, error) {
	if len(text) < 20 {
		return 0, sandbox.Deps.Std.Errorf("verb: %q is not a valid RFC3339 timestamp: too short", text)
	}
	if text[4] != '-' || text[7] != '-' || (text[10] != 'T' && text[10] != 't') || text[13] != ':' || text[16] != ':' {
		return 0, sandbox.Deps.Std.Errorf("verb: %q is not a valid RFC3339 timestamp: expected the layout %q", text, Layout)
	}

	year, yearOk := digits(sandbox, text[0:4])
	month, monthOk := digits(sandbox, text[5:7])
	day, dayOk := digits(sandbox, text[8:10])
	hour, hourOk := digits(sandbox, text[11:13])
	minute, minuteOk := digits(sandbox, text[14:16])
	second, secondOk := digits(sandbox, text[17:19])
	if !yearOk || !monthOk || !dayOk || !hourOk || !minuteOk || !secondOk {
		return 0, sandbox.Deps.Std.Errorf("verb: %q is not a valid RFC3339 timestamp: every field must be digits", text)
	}

	if month < 1 || month > 12 {
		return 0, sandbox.Deps.Std.Errorf("verb: %q is not a valid RFC3339 timestamp: month %d is out of range", text, month)
	}
	if day < 1 || day > daysInMonth(sandbox, year, month) {
		return 0, sandbox.Deps.Std.Errorf("verb: %q is not a valid RFC3339 timestamp: day %d is out of range for month %d", text, day, month)
	}
	if hour > 23 || minute > 59 || second > 59 {
		return 0, sandbox.Deps.Std.Errorf("verb: %q is not a valid RFC3339 timestamp: time %02d:%02d:%02d is out of range", text, hour, minute, second)
	}

	rest := text[19:]

	fraction := 0
	if len(rest) > 0 && rest[0] == '.' {
		read := 0
		for read+1 < len(rest) && isDigit(sandbox, rest[read+1]) {
			read++
		}
		if read == 0 {
			return 0, sandbox.Deps.Std.Errorf("verb: %q is not a valid RFC3339 timestamp: the fractional second has no digits", text)
		}
		// The grammar puts no ceiling on the fractional second, but the
		// instant is reported in nanoseconds, so a tenth digit and whatever
		// follows it is read and dropped rather than refused.
		kept := read
		if kept > 9 {
			kept = 9
		}
		value, _ := digits(sandbox, rest[1:kept+1])
		for scale := kept; scale < 9; scale++ {
			value *= 10
		}
		fraction = value
		rest = rest[read+1:]
	}

	offset, err := zoneOffset(sandbox, text, rest)
	if err != nil {
		return 0, err
	}

	days := daysFromCivil(sandbox, year, month, day)
	seconds := days*86400 + int64(hour)*3600 + int64(minute)*60 + int64(second) - offset
	if seconds > maxSeconds || seconds < -maxSeconds {
		return 0, sandbox.Deps.Std.Errorf("verb: %q is outside the range an int64 of nanoseconds can hold", text)
	}
	return seconds*nanosPerSecond + int64(fraction), nil
}

// zoneOffset reads the zone that closes an RFC 3339 timestamp — "Z" for UTC,
// or "+HH:MM" / "-HH:MM" — and returns it as a number of seconds to subtract
// from the wall clock to reach UTC. text is carried only so a failure can
// name the timestamp it came from.
func zoneOffset(sandbox *api.Sandbox, text string, zone string) (int64, error) {
	if zone == "Z" || zone == "z" {
		return 0, nil
	}
	if len(zone) != 6 || (zone[0] != '+' && zone[0] != '-') || zone[3] != ':' {
		return 0, sandbox.Deps.Std.Errorf("verb: %q is not a valid RFC3339 timestamp: the zone must be Z or ±HH:MM", text)
	}
	hours, hoursOk := digits(sandbox, zone[1:3])
	minutes, minutesOk := digits(sandbox, zone[4:6])
	if !hoursOk || !minutesOk || hours > 23 || minutes > 59 {
		return 0, sandbox.Deps.Std.Errorf("verb: %q is not a valid RFC3339 timestamp: the zone offset is out of range", text)
	}
	offset := int64(hours)*3600 + int64(minutes)*60
	if zone[0] == '-' {
		return -offset, nil
	}
	return offset, nil
}

// daysFromCivil returns the number of days between 1970-01-01 and the given
// civil date, negative before the epoch. It is the shift-the-year-to-March
// formula, which needs no table and no branch on the leap rule: with March as
// the first month the leap day lands at the end of the year, so the month
// lengths run in a fixed 153-day-per-five-month pattern.
func daysFromCivil(sandbox *api.Sandbox, year int, month int, day int) int64 {
	shiftedYear := int64(year)
	if month <= 2 {
		shiftedYear--
	}

	era := shiftedYear / 400
	if shiftedYear < 0 {
		era = (shiftedYear - 399) / 400
	}
	yearOfEra := shiftedYear - era*400

	shiftedMonth := int64(month) + 9
	if month > 2 {
		shiftedMonth = int64(month) - 3
	}
	dayOfYear := (153*shiftedMonth+2)/5 + int64(day) - 1
	dayOfEra := yearOfEra*365 + yearOfEra/4 - yearOfEra/100 + dayOfYear

	return era*146097 + dayOfEra - 719468
}

// daysInMonth returns how many days the given month of the given year holds.
func daysInMonth(sandbox *api.Sandbox, year int, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeapYear(sandbox, year) {
			return 29
		}
		return 28
	}
	return 0
}

// isLeapYear reports whether the given year holds a 29th of February, by the
// Gregorian rule.
func isLeapYear(sandbox *api.Sandbox, year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 != 0 {
		return true
	}
	return year%400 == 0
}

// digits reads text as an unsigned decimal number, reporting false when any
// character of it is not a digit. It stands in for strconv.Atoi, which would
// accept a sign and is reached through a dep besides — a fixed-width field of
// a timestamp is neither signed nor worth a dep call.
func digits(sandbox *api.Sandbox, text string) (int, bool) {
	if len(text) == 0 {
		return 0, false
	}
	value := 0
	for index := 0; index < len(text); index++ {
		if !isDigit(sandbox, text[index]) {
			return 0, false
		}
		value = value*10 + int(text[index]-'0')
	}
	return value, true
}

// isDigit reports whether character is one of '0' to '9'.
func isDigit(sandbox *api.Sandbox, character byte) bool {
	return character >= '0' && character <= '9'
}
