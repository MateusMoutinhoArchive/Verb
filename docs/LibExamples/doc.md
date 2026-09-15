# LibExamples

Every example of Verb used as a Go module. Each one is a `package main` program that
runs with its own directory as the working directory and writes only into its own `TestDir`,
so it can be read as documentation and copied as a starting point. It ends by copying out of
`TestDir` into `AssertDir` the paths it asserts — `os.CopyFS(dst, os.DirFS(src))`, one call per
path, each keeping the place it holds in the tree.

`agnos exec-test` runs them all and checks each against the `result.yaml` beside it — the
golden holding the output, the exit code and the sha256 of every `AssertDir` file, written by
`exec-test` and never by hand. [Workflow](../Workflow/doc.md) has the commands that add and
remove one.

| Example | Description | Source |
|---|---|---|
| `key-values` | Read key=value arguments by prefix, and see an empty value refused | [example.go](../../examples/lib/key-values/example.go) |
| `options` | Read a repeatable --flag value option, every occurrence of it, and a missing value refused | [example.go](../../examples/lib/options/example.go) |
| `positional-args` | Read an argument by its absolute place on the command line, and an out-of-range index refused | [example.go](../../examples/lib/positional-args/example.go) |
| `presence` | Read a boolean flag by any of its spellings, and see it consumed once | [example.go](../../examples/lib/presence/example.go) |
| `timestamps` | Read RFC 3339 timestamps as nanoseconds since the epoch — zones, fractions, the leap day and what is refused | [example.go](../../examples/lib/timestamps/example.go) |
| `typed-values` | Read an int, a float and an RFC 3339 timestamp, and see a malformed value refused while still counting as read | [example.go](../../examples/lib/typed-values/example.go) |
| `unused-mechanic` | Read every declared flag and option first, then drain whatever positional arguments are left over | [example.go](../../examples/lib/unused-mechanic/example.go) |

