# Verb

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


## Documentation

### LibUsage

Using the project as a Go module - wiring the deps, calling the sandbox

| Doc | Description |
| --- | --- |
| [LibUsage](docs/LibUsage/doc.md) | Use Verb as a Go module: wire the deps, build the sandbox, call its API |
| [PublicApi](docs/PublicApi/doc.md) | Every exported symbol of Verb, generated from the contract sources and their doc comments |
| [LibExamples](docs/LibExamples/doc.md) | Index of every runnable example of Verb as a Go module |
| [Embed Verb](docs/EmbedVerb/doc.md) | Taking Verb as a dep of another agnos project, or restating its shape in a contract of your own |
| [Timestamps](docs/Timestamps/doc.md) | The RFC 3339 grammar the Timestamp getters accept, and why an instant crosses as an int64 |
| [Unused Mechanic](docs/UnusedMechanic/doc.md) | Every argument starts unread: how a matched argument is marked and what is left over |

### Architecture

How the project is put together - layers, boundaries, data flow

| Doc | Description |
| --- | --- |
| [Adapters](docs/Adapters/doc.md) | Contract, adapter and available: three units, one field of Deps, and who fills it |
| [Embed Verb](docs/EmbedVerb/doc.md) | Taking Verb as a dep of another agnos project, or restating its shape in a contract of your own |
| [Unused Mechanic](docs/UnusedMechanic/doc.md) | Every argument starts unread: how a matched argument is marked and what is left over |

### Development

Changing this repository - schema, build mechanics, recipes

| Doc | Description |
| --- | --- |
| [Requirements](docs/Requirements/doc.md) | The two tools this project needs — Go and agnos — installed per platform |
| [Workflow](docs/Workflow/doc.md) | Every change this project takes and the agnos command that makes it |
| [Rules](docs/Rules/doc.md) | Every rule the generators, `verify` and the hand-written files must hold to |
| [Structure](docs/Structure/doc.md) | The project schema: what lives where, what is generated, what verify enforces |

### Reference

Lookup tables - schemas, file formats, generated file listings

| Doc | Description |
| --- | --- |
| [EntriesYaml](docs/EntriesYaml/doc.md) | Every key of a command's entries.yaml and what the generated code does with it |
| [DepList](docs/DepList/doc.md) | Every dep `agnos add-dep` can add, the adapters that fill it, and what backs each one |
| [GeneratedFiles](docs/GeneratedFiles/doc.md) | Every file agnos writes into this project and whether build overwrites it |
| [LibExamples](docs/LibExamples/doc.md) | Index of every runnable example of Verb as a Go module |
| [Timestamps](docs/Timestamps/doc.md) | The RFC 3339 grammar the Timestamp getters accept, and why an instant crosses as an int64 |

## License

This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <https://unlicense.org>

