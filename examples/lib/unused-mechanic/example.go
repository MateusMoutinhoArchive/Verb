package main

import (
	"fmt"
	"os"

	"github.com/MateusMoutinhoOrg/Verb/adapters/availables/standard"
	"github.com/MateusMoutinhoOrg/Verb/sandbox"
)

// The Unused Mechanic: a whole command line read without a grammar.
//
// Every argument starts out unread, and every reader marks what it matched.
// So a program declares what it knows — its flags and its options — and
// whatever is still unread afterwards is exactly the positional arguments
// nobody asked for, in the order the caller typed them. Nothing has to know
// in advance how many there are, or where they sit between the flags.

func main() {

	deps := standard.New()    // std, stringsdeps
	lib := sandbox.New(&deps) // *api.Sandbox

	parser := lib.Argv.New([]string{
		"first.txt", "--quiet", "--output", "build/", "second.txt", "-f", "third.txt",
	})

	// 1. Everything this program declares, in whatever order suits it.
	quiet := parser.IsPresent([]string{"-q", "--quiet"})
	force := parser.IsPresent([]string{"-f", "--force"})

	output, err := parser.GetStringOption([]string{"-o", "--output"}, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("quiet:", quiet)
	fmt.Println("force:", force)
	fmt.Println("output:", output)

	// 2. Whatever is left. The loop ends on the error the parser returns when
	// nothing is unread, so it needs no count and no index of its own.
	var files []string
	for {
		file, err := parser.GetNextStringArg()
		if err != nil {
			fmt.Println("drained:", err)
			break
		}
		files = append(files, file)
		fmt.Println("file:", file)
	}

	writeOut("unused-mechanic.txt", fmt.Sprintf("quiet=%t force=%t output=%s files=%v\n", quiet, force, output, files))
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
