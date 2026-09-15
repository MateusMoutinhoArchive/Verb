package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MateusMoutinhoOrg/Verb/adapters/availables/standard"
	"github.com/MateusMoutinhoOrg/Verb/sandbox"
)

// What a Timestamp getter accepts, and what it hands back.
//
// An instant leaves Verb as an int64 of nanoseconds since the Unix epoch,
// UTC, never as a time.Time: the sandbox may not name `time`, and an api type
// has to be convertible for a consumer to copy this contract into a deps
// package of its own. A caller that wants a time.Time writes
// `time.Unix(0, nanos).UTC()` on its own side — one line, outside the library.
//
// The grammar is RFC 3339 itself, implemented in arithmetic: the zone offset
// is applied, the calendar is real, and a fractional second past the ninth
// digit is dropped rather than refused.

// Accepted is one timestamp per thing worth pinning: the epoch itself, the
// second before it, both signs of zone offset, a fractional second, a leap
// day, and a fraction longer than nanoseconds can hold.
var Accepted = []string{
	"1970-01-01T00:00:00Z",
	"1969-12-31T23:59:59Z",
	"2024-01-02T15:04:05Z",
	"2024-01-02T12:04:05-03:00",
	"2024-01-02T18:34:05+03:30",
	"2024-01-02T15:04:05.25Z",
	"2024-02-29T00:00:00Z",
	"2024-01-02T15:04:05.1234567891Z",
}

// Refused is one timestamp per way the grammar or the calendar says no.
var Refused = []string{
	"2023-02-29T00:00:00Z",
	"2024-04-31T00:00:00Z",
	"2024-01-02T24:00:00Z",
	"2024-01-02T15:04:05+0300",
	"2024-01-02 15:04:05Z",
	"2024-01-02T15:04:05",
	"1500-01-01T00:00:00Z",
}

func main() {

	deps := standard.New()    // std, stringsdeps
	lib := sandbox.New(&deps) // *api.Sandbox

	var report strings.Builder

	for _, text := range Accepted {
		// Each timestamp is read as the value of "--at", so this walks the
		// same path a real command line takes.
		parser := lib.Argv.New([]string{"--at", text})
		nanos, err := parser.GetTimestampOption([]string{"--at"}, 0)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%-32s %d\n", text, nanos)
		fmt.Fprintf(&report, "%s=%d\n", text, nanos)
	}

	// The third and fourth entries of Accepted are one instant written in
	// three zones, so all three come back as the same number.
	fmt.Println()

	for _, text := range Refused {
		parser := lib.Argv.New([]string{"--at", text})
		_, err := parser.GetTimestampOption([]string{"--at"}, 0)
		if err == nil {
			panic(fmt.Sprintf("%q should have been refused", text))
		}
		fmt.Println("refused:", err)

		// The argument was found and read, so it is marked used even though
		// the value turned out malformed.
		fmt.Fprintf(&report, "%s=refused used=%v\n", text, parser.Used)
	}

	writeOut("timestamps.txt", report.String())
}

// writeOut records what the example parsed under TestDir and copies it into
// AssertDir, which is what the golden result.yaml holds.
func writeOut(name string, content string) {
	if err := os.MkdirAll("TestDir", 0o755); err != nil {
		panic(err)
	}
	if err := os.WriteFile("TestDir/"+name, []byte(content), 0o644); err != nil {
		panic(err)
	}
	if err := os.CopyFS("AssertDir", os.DirFS("TestDir")); err != nil {
		panic(err)
	}
}
