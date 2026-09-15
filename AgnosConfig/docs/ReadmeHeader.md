# {{.Name}}

[![Go Reference](https://pkg.go.dev/badge/github.com/MateusMoutinhoOrg/Verb.svg)](https://pkg.go.dev/github.com/MateusMoutinhoOrg/Verb)
[![Release](https://img.shields.io/github/v/release/MateusMoutinhoOrg/Verb)](https://github.com/MateusMoutinhoOrg/Verb/releases/latest)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.25-blue)](go.mod)

**An OS-independent argv parser.** Flags, options, `key=value` arguments and positionals,
read out of any `[]string` — and whatever nothing asked for, handed back in order. No
globals, no registration pass, no `os` anywhere in the library.

```bash
go get github.com/MateusMoutinhoOrg/Verb@latest
```

```go
package main

import (
	"fmt"
	"os"

	"github.com/MateusMoutinhoOrg/Verb/adapters/availables/standard"
	"github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	deps := standard.New()    // std, stringsdeps
	lib := sandbox.New(&deps) // *api.Sandbox

	parser := lib.Argv.New(os.Args[1:])

	// 1. Everything this program declares, in whatever order suits it.
	quiet := parser.IsPresent([]string{"-q", "--quiet"})
	output, err := parser.GetStringOption([]string{"-o", "--output"}, 0)
	if err != nil {
		output = "."
	}

	// 2. Whatever is left over is the positional arguments — no count, no
	//    index, no grammar declared up front.
	for {
		file, err := parser.GetNextStringArg()
		if err != nil {
			break
		}
		fmt.Println(quiet, output, file)
	}
}
```

Every argument starts unread, and every reader marks what it matched. So a program asks for
what it knows about first, and `GetNextStringArg` then drains exactly the arguments nobody
claimed — that is the [Unused Mechanic](docs/UnusedMechanic/doc.md), and it is why Verb
needs no grammar and no ordering rules.

Each reader family — `IsPresent`, `Option`, `Arg`, `NextArg`, `KeyValues` — comes in four
types: `String`, `Int`, `Double` and `Timestamp`. A timestamp is RFC 3339 in and an `int64`
of nanoseconds out, parsed in arithmetic alone: the library names neither `time` nor `os`,
which is what lets any project copy its contract and fill it by assigning fields across.

| Start here | For |
|---|---|
| [LibUsage](docs/LibUsage/doc.md) | wiring the deps and calling the sandbox |
| [PublicApi](docs/PublicApi/doc.md) | every exported symbol, generated from the contracts |
| [LibExamples](docs/LibExamples/doc.md) | seven runnable programs, each checked against a golden |
| [UnusedMechanic](docs/UnusedMechanic/doc.md) | who marks what used, and in what order to read |
| [EmbedVerb](docs/EmbedVerb/doc.md) | taking Verb as a dep, or restating its shape yourself |

This repository is generated and checked by [agnos](https://github.com/MateusMoutinhoOrg/Agnos):
`agnos build` rewrites every generated file, `agnos verify` checks the schema, and
`agnos exec-test` runs every example against its golden. See
[Requirements](docs/Requirements/doc.md) and [Workflow](docs/Workflow/doc.md).
