# UnusedMechanic

Every argument of an `api.Parser` starts unread. `Used` tracks, index for index against
`Args`, what has been consumed. Whatever is still unread after a program has asked for
everything it declares is exactly the positional arguments nobody asked for — so a command
line is read without a grammar, and without a pass that has to run before the others.

## Who marks what

| Function | Marks |
|---|---|
| `IsPresent` | the first **unused** argument matching one of the flags, when there is one |
| `GetStringOption` and its typed variants | the matched flag **and** the argument after it |
| `GetStringArg` and its typed variants | the index given |
| `GetNextStringArg` and its typed variants | the first still-unused argument |
| `GetStringKeyValues` and its typed variants | the matched argument |
| `GetOptionsSize`, `GetKeyValuesSize` | nothing — they only count |

A typed getter marks its match used **even when the parse then fails**: the argument was
found and read, only its value turned out malformed. That is what keeps a bad value out of
the leftovers.

## Occurrence numbering is not a cursor

`GetOptionsSize` and `GetKeyValuesSize` count every match, used or not, and the occurrence
index of a getter counts the same way. So `0..size-1` is a stable range that does not shift
as the loop consumes it, and re-reading occurrence 0 returns the same argument again.
`IsPresent` is the one reader that skips what is already used, because a flag carries no
value and a second call asking the same question means a second flag.

## The order a program reads in

1. Every flag, with `IsPresent`.
2. Every option and key/value, with `GetOptionsSize` / `GetKeyValuesSize` and the getter.
3. Whatever is left, with `GetNextStringArg` until it errors.

Step 3 needs no count and no index: the error that ends the loop is the parser reporting
that nothing is unread. Reversing the order breaks it — `GetNextStringArg` called first
returns the first argument whatever it is, flag included.

Worked through end to end in [`examples/lib/unused-mechanic`](../../examples/lib/unused-mechanic/example.go).

## Args and Used are read-only

Both are exported so the closures filled by `sandbox/internal/argv` can reach them across
packages, and a caller may read `Used` to see what a parse consumed. Writing to either
leaves them out of sync and makes matching undefined; neither slice is ever regrown, which
is what makes the parser value `Argv.New` returns share its read state with the closures
inside it.
