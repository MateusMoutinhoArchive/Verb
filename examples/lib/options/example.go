package main

import (
	"fmt"
	"os"

	"github.com/MateusMoutinhoOrg/Verb/adapters/availables/standard"
	"github.com/MateusMoutinhoOrg/Verb/sandbox"
)

// An option is a flag followed by its value: `--username mateus`.
//
// Options repeat. GetOptionsSize counts the occurrences without consuming
// anything, so the loop below is written against a number the parser reports
// rather than against a guess — and occurrence numbering counts every match,
// so it does not shift as the loop consumes them.

// Spellings is every way this program accepts the option being written.
var Spellings = []string{"-username", "--username"}

func main() {

	deps := standard.New()    // std, stringsdeps
	lib := sandbox.New(&deps) // *api.Sandbox

	parser := lib.Argv.New([]string{"--username", "mateus", "-username", "ana", "--username"})

	size := parser.GetOptionsSize(Spellings)
	fmt.Println("occurrences:", size)

	for occurrence := 0; occurrence < size; occurrence++ {
		// The third occurrence is the last argument of the argv, so it has
		// nothing after it to read as a value: that is an error, not a panic
		// and not an empty string.
		username, err := parser.GetStringOption(Spellings, occurrence)
		if err != nil {
			fmt.Println("refused:", err)
			continue
		}
		fmt.Printf("username %d: %s\n", occurrence, username)
	}

	// An occurrence past the last one is refused the same way.
	if _, err := parser.GetStringOption(Spellings, size); err != nil {
		fmt.Println("refused:", err)
	}

	writeOut("options.txt", fmt.Sprintf("occurrences=%d used=%v\n", size, parser.Used))
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
