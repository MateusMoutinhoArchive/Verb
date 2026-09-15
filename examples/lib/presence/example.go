package main

import (
	"fmt"
	"os"

	"github.com/MateusMoutinhoOrg/Verb/adapters/availables/standard"
	"github.com/MateusMoutinhoOrg/Verb/sandbox"
)

// Reading a boolean flag, whichever way the caller spelled it.
//
// IsPresent takes every spelling of one flag at once, so "-f", "--f", "-file"
// and "--file" are one question and not four. It is the one reader that
// returns no error: a flag that is absent is an answer, not a failure.

func main() {

	deps := standard.New()    // std, stringsdeps
	lib := sandbox.New(&deps) // *api.Sandbox

	// A real program passes os.Args[1:] here. An example passes the argv it
	// is about, so it reads as documentation and runs the same everywhere.
	parser := lib.Argv.New([]string{"--file", "report.txt", "-q"})

	force := parser.IsPresent([]string{"-f", "--f", "-file", "--file"})
	quiet := parser.IsPresent([]string{"-q", "--quiet"})
	verbose := parser.IsPresent([]string{"-v", "--verbose"})

	fmt.Println("file:", force)
	fmt.Println("quiet:", quiet)
	fmt.Println("verbose:", verbose)

	// A matched flag is marked used, so the argv it came from is left with
	// only what nothing has asked for yet — see the unused-mechanic example.
	fmt.Println("used:", parser.Used)

	report := fmt.Sprintf("file=%t quiet=%t verbose=%t used=%v\n", force, quiet, verbose, parser.Used)
	writeOut("presence.txt", report)
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
