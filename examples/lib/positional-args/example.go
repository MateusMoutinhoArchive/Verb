package main

import (
	"fmt"
	"os"

	"github.com/MateusMoutinhoOrg/Verb/adapters/availables/standard"
	"github.com/MateusMoutinhoOrg/Verb/sandbox"
)

// Reading an argument by where it stands.
//
// GetStringArg numbers the argv the way the command line does: index 0 is the
// first argument after the program name, whatever it happens to be and
// whatever has already been read. It is the reader to use when a program's
// shape is fixed — `app <command> <target>` — rather than flag-driven.

func main() {

	deps := standard.New()    // std, stringsdeps
	lib := sandbox.New(&deps) // *api.Sandbox

	parser := lib.Argv.New([]string{"build", "--quiet", "report.txt"})

	command, err := parser.GetStringArg(0)
	if err != nil {
		panic(err)
	}
	fmt.Println("command:", command)

	target, err := parser.GetStringArg(2)
	if err != nil {
		panic(err)
	}
	fmt.Println("target:", target)

	// Index 1 was never asked for, so it stays unread — position numbering is
	// absolute, not a cursor.
	fmt.Println("used:", parser.Used)

	// An index past the end is refused.
	if _, err := parser.GetStringArg(9); err != nil {
		fmt.Println("refused:", err)
	}

	writeOut("positional-args.txt", fmt.Sprintf("command=%s target=%s used=%v\n", command, target, parser.Used))
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
