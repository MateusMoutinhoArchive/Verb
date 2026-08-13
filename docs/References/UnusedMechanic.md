# The Unused Mechanic

## Description
Explains how the library tracks which arguments have already been read, and how that tracking turns "everything I did not explicitly ask for" into a list a program can drain in order — the mechanic behind `api.Lib.Used` and the `GetNext*Arg` family.

---

## What Used Tracks

`lib.New` stores the argument vector on [`api.Lib.Args`](/docs/References/PublicApi/api.Args.md) and allocates [`api.Lib.Used`](/docs/References/PublicApi/api.Used.md), a `[]bool` of the same length. Index for index, `Used[i]` answers one question: has `Args[i]` already been consumed by a call?

Every argument starts unread — `Used` is entirely `false` — and it only ever grows more true over the `Lib`'s lifetime. Nothing resets it, so the order in which a program reads its arguments is what decides what is left at the end.

---

## Who Marks What

| Call | What it marks used |
|------|--------------------|
| `IsPresent` | the single argument that matched one of the flags |
| `GetStringOption` and its typed variants | both the flag **and** the value following it |
| `GetStringArg` and its typed variants | the argument at the requested index |
| `GetStringKeyValues` and its typed variants | the matched `key=value` argument |
| `GetNextStringArg` and its typed variants | the first argument still unread |
| `GetOptionsSize`, `GetKeyValuesSize` | **nothing** — they only count |

A typed getter marks its argument used even when parsing then fails: the argument was found and read, only its value turned out malformed. The two `*Size` functions are the exception on purpose — counting occurrences is what a caller does *before* looping over them, so it must not consume anything.

---

## Draining the Leftovers

Because every expected flag and option consumes itself as it is read, whatever remains unread is exactly the positional arguments the program never explicitly asked for. `GetNextStringArg` walks `Used` in order and returns the first one still `false`, marking it used, so repeated calls drain them left to right — and return an error once nothing is left.

For a command line like:

```bash
./cli -o teste.out --quiet teste.c
```

reading the expected options first leaves the trailing filename as the only unread argument:

```go
package main

import (
	"os"

	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// Build the lib directly from the real process argv.
	lib := verblib.New(os.Args[1:])

	// marks -o and teste.out as used
	output, err := lib.GetStringOption([]string{"-o", "--o", "--output", "--out"}, 0)
	if err != nil {
		panic(err)
	}

	// marks --quiet as used
	quiet := lib.IsPresent([]string{"-q", "--q", "--quiet"})

	// gets teste.c, the first argument still unused
	file, err := lib.GetNextStringArg()
	if err != nil {
		panic(err)
	}

	println(output, quiet, file)
}
```

Reading the leftovers *before* the options would hand `-o` back as the "next" argument, so a program that mixes both always reads its flags and options first. Walking through that order step by step is [UseUnusedMechanic.md](/docs/Tutorials/UseUnusedMechanic.md).
