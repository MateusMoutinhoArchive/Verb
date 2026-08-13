// Package argv holds the matching and parsing helpers every public function
// of the library is built from: locating an argument in api.Lib.Args, marking
// it used, and turning the raw text it holds into a typed value. It declares
// no types and no factories — the packages under sandbox/lib/ call it.
package argv

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/MateusMoutinhoOrg/Verb/sandbox/config"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
)

// MatchesFlag reports whether arg equals one of the given flag spellings.
func MatchesFlag(arg string, flags []string) bool {
	for _, f := range flags {
		if arg == f {
			return true
		}
	}
	return false
}

// MatchPrefix reports whether arg starts with one of the given key=value
// prefixes, returning the matched prefix.
func MatchPrefix(arg string, prefixes []string) (string, bool) {
	for _, p := range prefixes {
		if strings.HasPrefix(arg, p) {
			return p, true
		}
	}
	return "", false
}

// CountFlags counts how many arguments equal one of flags, without touching
// l.Used.
func CountFlags(l *api.Lib, flags []string) int {
	count := 0
	for _, a := range l.Args {
		if MatchesFlag(a, flags) {
			count++
		}
	}
	return count
}

// CountPrefixes counts how many arguments start with one of prefixes, without
// touching l.Used.
func CountPrefixes(l *api.Lib, prefixes []string) int {
	count := 0
	for _, a := range l.Args {
		if _, ok := MatchPrefix(a, prefixes); ok {
			count++
		}
	}
	return count
}

// FirstUnusedFlag returns the index of the first not-yet-used argument
// matching one of flags, marking it used.
func FirstUnusedFlag(l *api.Lib, flags []string) (int, bool) {
	for i, a := range l.Args {
		if l.Used[i] {
			continue
		}
		if MatchesFlag(a, flags) {
			l.Used[i] = true
			return i, true
		}
	}
	return -1, false
}

// findOptionIndex returns the index in l.Args of the occurrence-th argument
// matching one of flags, regardless of Used.
func findOptionIndex(l *api.Lib, flags []string, occurrence int) (int, error) {
	count := 0
	for i, a := range l.Args {
		if MatchesFlag(a, flags) {
			if count == occurrence {
				return i, nil
			}
			count++
		}
	}
	return -1, fmt.Errorf("verb: option %v: occurrence %d not found (only %d present)", flags, occurrence, count)
}

// OptionValue backs the Option getters: it locates the occurrence-th flag,
// marks it and its following value as used, and returns that value.
func OptionValue(l *api.Lib, flags []string, occurrence int) (string, error) {
	idx, err := findOptionIndex(l, flags, occurrence)
	if err != nil {
		return "", err
	}
	if idx+1 >= len(l.Args) {
		return "", fmt.Errorf("verb: option %v: flag %q has no following value", flags, l.Args[idx])
	}
	l.Used[idx] = true
	l.Used[idx+1] = true
	return l.Args[idx+1], nil
}

// ArgValue backs the Arg getters: it validates index against l.Args, marks it
// used, and returns the argument at that absolute position.
func ArgValue(l *api.Lib, index int) (string, error) {
	if index < 0 || index >= len(l.Args) {
		return "", fmt.Errorf("verb: arg index %d out of range (have %d arguments)", index, len(l.Args))
	}
	l.Used[index] = true
	return l.Args[index], nil
}

// NextArgValue backs the NextArg getters — the Unused Mechanic: it finds the
// first not-yet-used argument in order, marks it used, and returns it.
func NextArgValue(l *api.Lib) (string, error) {
	for i, used := range l.Used {
		if !used {
			l.Used[i] = true
			return l.Args[i], nil
		}
	}
	return "", fmt.Errorf("verb: no unused arguments remaining")
}

// KeyValuesValue backs the KeyValues getters: it locates the occurrence-th
// argument starting with one of prefixes, marks it used, and returns the text
// after the matched prefix.
func KeyValuesValue(l *api.Lib, prefixes []string, occurrence int) (string, error) {
	count := 0
	for i, a := range l.Args {
		prefix, ok := MatchPrefix(a, prefixes)
		if !ok {
			continue
		}
		if count == occurrence {
			l.Used[i] = true
			value := a[len(prefix):]
			if value == "" {
				return "", fmt.Errorf("verb: key/value %v: occurrence %d has an empty value", prefixes, occurrence)
			}
			return value, nil
		}
		count++
	}
	return "", fmt.Errorf("verb: key/value %v: occurrence %d not found (only %d present)", prefixes, occurrence, count)
}

// ParseInt parses s as a base-10 integer for the Int-typed getters.
func ParseInt(s string) (int, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("verb: %q is not a valid integer: %w", s, err)
	}
	return v, nil
}

// ParseDouble parses s as a 64-bit float for the Double-typed getters.
func ParseDouble(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("verb: %q is not a valid number: %w", s, err)
	}
	return v, nil
}

// ParseTimestamp parses s with config.TimestampLayout for the Timestamp-typed
// getters.
func ParseTimestamp(s string) (time.Time, error) {
	v, err := time.Parse(config.TimestampLayout, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("verb: %q is not a valid RFC3339 timestamp: %w", s, err)
	}
	return v, nil
}
