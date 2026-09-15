# EmbedVerb

Verb is a closed sandbox with no OS-bound code in it, so a program embeds it the way it
embeds any agnos library — and an agnos project can install it as a dep, contract and shim
generated for it.

## As a plain Go module

```bash
go get github.com/MateusMoutinhoOrg/Verb@latest
```

```go
deps := standard.New()    // std, stringsdeps
lib := sandbox.New(&deps) // *api.Sandbox

parser := lib.Argv.New(os.Args[1:])
```

`standard.New()` is the whole of what Verb asks of the outside world: formatting and text
conversion. Wiring, patching and the list of contracts are in [LibUsage](../LibUsage/doc.md).

## As a dep of another agnos project

```bash
agnos add-dep github.com/MateusMoutinhoOrg/Verb@latest --as argvdeps --remote-available standard
```

That writes two files and nothing else ([GeneratedFiles](../GeneratedFiles/doc.md)):

| File | What it is |
|---|---|
| `sandbox/deps/argvdeps/*.go` | a copy of Verb's `sandbox/api/`, so the consuming sandbox names its own type and imports nothing |
| `adapters/libs/argvdeps/argvdeps.go` | the shim: `Bind` builds a Verb sandbox and assigns its fields across |

The consuming sandbox then reads `sandbox.Deps.Argvdeps.New(args)` and never imports Verb.
`set-dep` moves the copy to another version of the module.

This is why `api.Parser` holds function fields and not methods, and why a timestamp crosses
as an `int64` rather than a `time.Time` ([Timestamps](../Timestamps/doc.md)): the copy has
to be a *convertible* restatement of the contract, which
[Rules](../Rules/doc.md#layers) requires of every api type. An adapter fills the copy by
assigning the real fields straight into it — no wrapper type, no interface, no reflection.

## Restating the contract by hand

A project that is not an agnos repo can do the same thing by hand: declare a struct with
the fields it uses, and assign. Nothing in `api.Parser` names a type from another package,
so the restatement compiles against no import at all.

```go
// In your own package, naming only the fields you call.
type ArgvParser struct {
	IsPresent        func(flags []string) bool
	GetStringOption  func(flags []string, occurrence int) (string, error)
	GetNextStringArg func() (string, error)
}

parser := lib.Argv.New(os.Args[1:])
mine := ArgvParser{
	IsPresent:        parser.IsPresent,
	GetStringOption:  parser.GetStringOption,
	GetNextStringArg: parser.GetNextStringArg,
}
```

The `Used` state lives in the parser the fields were taken from, so a copy narrowed this way
still shares it: reading a flag through `mine` keeps that argument out of what
`parser.GetNextStringArg` later drains ([UnusedMechanic](../UnusedMechanic/doc.md)).

## A note on the catalog's `argvdeps`

`agnos list-deps` also ships a hand-written `argvdeps` contract with a `verb` adapter
pinned to an older release, from before this repository was an agnos project. It binds
`verblib.New(args)` directly. The generated route above — `add-dep <module>` — is the one
that follows this repository's api.
