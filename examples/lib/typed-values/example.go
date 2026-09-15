package main

import (
	"fmt"
	"os"

	"github.com/MateusMoutinhoOrg/Verb/adapters/availables/standard"
	"github.com/MateusMoutinhoOrg/Verb/sandbox"
)

// Every reader family comes in four types.
//
// String hands back the raw text; Int parses base-10; Double parses a
// float64; Timestamp parses RFC 3339 and reports nanoseconds since the Unix
// epoch, UTC — the sandbox may not name a time.Time, and an api type has to
// be convertible for a consumer to copy this contract into its own deps.
//
// A typed reader marks its match used even when the parse then fails: the
// argument was found and read, only its value turned out malformed. That is
// what keeps a bad value out of the leftovers the Unused Mechanic drains.

func main() {

	deps := standard.New()    // std, stringsdeps
	lib := sandbox.New(&deps) // *api.Sandbox

	parser := lib.Argv.New([]string{
		"--retries", "3",
		"--ratio", "0.25",
		"--since", "2024-01-02T15:04:05Z",
		"--until", "2024-01-02T12:04:05-03:00",
		"--port", "eighty",
		"--born", "2023-02-29T00:00:00Z",
		"report.txt",
	})

	retries, err := parser.GetIntOption([]string{"--retries"}, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("retries:", retries)

	ratio, err := parser.GetDoubleOption([]string{"--ratio"}, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("ratio:", ratio)

	// The two timestamps below are the same instant written in two zones, so
	// they come back as the same number: the offset is applied on the way in.
	since, err := parser.GetTimestampOption([]string{"--since"}, 0)
	if err != nil {
		panic(err)
	}
	until, err := parser.GetTimestampOption([]string{"--until"}, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("since:", since)
	fmt.Println("until:", until)
	fmt.Println("same instant:", since == until)

	// "eighty" is not a number. The read still happened, so --port and its
	// value are marked used and neither reaches the leftovers.
	if _, err := parser.GetIntOption([]string{"--port"}, 0); err != nil {
		fmt.Println("refused:", err)
	}

	// A date that does not exist is refused by the calendar, not by a format
	// check: 2023 is not a leap year, so it has no 29th of February.
	if _, err := parser.GetTimestampOption([]string{"--born"}, 0); err != nil {
		fmt.Println("refused:", err)
	}

	// Every reader above marked its match used, malformed values included, so
	// the one argument left over is the positional one nothing asked for.
	leftover, err := parser.GetNextStringArg()
	if err != nil {
		panic(err)
	}
	fmt.Println("leftover:", leftover)

	writeOut("typed-values.txt", fmt.Sprintf("retries=%d ratio=%v since=%d until=%d\n", retries, ratio, since, until))
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
