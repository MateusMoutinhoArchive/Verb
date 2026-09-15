package main

import (
	"fmt"
	"os"

	"github.com/MateusMoutinhoOrg/Verb/adapters/availables/standard"
	"github.com/MateusMoutinhoOrg/Verb/sandbox"
)

// A key/value is one argument carrying both halves: `username=mateus`.
//
// The separator is part of the prefix the reader is given, so the same
// mechanic reads `user:mateus` or `user/mateus` — Verb matches the text it is
// handed and never assumes a separator of its own.

// Prefixes is every way this program accepts the key being written. The "="
// belongs to the prefix.
var Prefixes = []string{"username=", "user="}

func main() {

	deps := standard.New()    // std, stringsdeps
	lib := sandbox.New(&deps) // *api.Sandbox

	parser := lib.Argv.New([]string{"username=mateus", "user=ana", "--verbose", "username="})

	size := parser.GetKeyValuesSize(Prefixes)
	fmt.Println("occurrences:", size)

	for occurrence := 0; occurrence < size; occurrence++ {
		// The last one carries the prefix and nothing else. A key written
		// with no value is refused rather than reported as an empty string,
		// which a caller would have to tell apart from a value of "".
		username, err := parser.GetStringKeyValues(Prefixes, occurrence)
		if err != nil {
			fmt.Println("refused:", err)
			continue
		}
		fmt.Printf("username %d: %s\n", occurrence, username)
	}

	writeOut("key-values.txt", fmt.Sprintf("occurrences=%d used=%v\n", size, parser.Used))
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
