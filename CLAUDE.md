# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

Verb is a **Go library template** demonstrating a struct-of-functions public API instead of interfaces. The value is in the *structure and documentation conventions*, not the (deliberately trivial) example code — an OS-independent argv parser (`IsPresent`, `GetStringOption`/`GetIntOption`/…, `GetStringArg`/…, `GetNextStringArg`/… "Unused Mechanic", `GetStringKeyValues`/…). This is not a placeholder to swap out casually: adapting the template into a real library is itself a documented, multi-file procedure (see the table below), not a simple find-and-replace.

The repository follows the same conventions as its sibling template [Agnos-Lib](https://github.com/mateusmoutinho/Agnos-Lib) — theme indexes, `docs/References/Specs/`, `sandbox/lib/publicfunctions/`, `examples/libraryExamples/` — minus the dependency-injection layer, which Verb does not need: its only input is the argument vector, passed as a plain argument.

## Commands

```bash
go build ./...                                                  # build everything
go run ./examples/libraryExamples/Presence/Presence.go          # run an example (also: Options, KeyValues, StringArg, NextArg)
go test ./...                                                   # run tests (none exist yet)
```

There is no lint config, CI, or test suite in the repo yet. When adding lib functions, `go build ./...` is the primary verification step.

## Architecture

Two top-level code trees, wired through **structs of function fields** (never interfaces), with a strict one-way dependency flow:

```
sandbox/  ◀──  examples/libraryExamples/
(closed)       (consume the lib)

lib.New(args)  ──▶  api.Lib
(entry point)       (the only output struct today, filled by sandbox/lib factories)
```

Contracts are structs whose fields hold functions, and **every** one of them is filled by **factories** — `func <Field>Factory(carrier *T) <FieldType>` bodies that return one closure reading the carrier at call time, with the assignment made explicitly by the caller. The carrier is the `api` struct, and `New` assigns the result (`l.IsPresent = publicfunctions.IsPresentFactory(&l)`, reading `l.Args`/`l.Used` inside the closure). No methods bound into fields, no internal mirror type. This is a binding rule — see the `Factories` spec (`docs/References/Specs/Factories/`) and `docs/References/StructContracts.md`.

One trade-off, not caught by the compiler: **completeness is unchecked** — a field no factory fills is nil and panics on first call, so every factory must be called from `New`.

`sandbox/` is a **closed sandbox**: nothing in it may import `examples/`, a third-party module, or an OS-bound stdlib package (`os`, `net`, `syscall`, …). Every input the library needs from the outside world arrives as a plain function argument, e.g. `lib.New(args []string)`. This is a binding rule — see `docs/References/SandboxIsolation.md`.

- **`sandbox/new.go`** — package `lib`, the only wiring point consumers touch: `New(args []string) api.Lib`. Never imports anything OS-bound. Importers alias it: `verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"`. It just delegates to `sandbox/lib.New`.
- **`sandbox/contracts/api/api.go`** — the output structs the lib hands back (currently just `Lib`). Every field must be exported, or `sandbox/lib/` cannot fill it. **Every type in the project is declared here**, never under `sandbox/lib/`. Types only — no function bodies.
- **`sandbox/config/`** — compile-time constants the sandbox reads (`TimestampLayout`). Nothing outside the sandbox imports it.
- **`sandbox/lib/new.go`** — `New(args []string) api.Lib`: the factory aggregate, assigning every `publicfunctions` factory's return value into its field.
- **`sandbox/lib/publicfunctions/<Field>.go`** — one file per function field of `api.Lib`, holding that field's `<Field>Factory(l *api.Lib)`.
- **`sandbox/lib/argv/argv.go`** — shared matching/consuming/parsing helpers the public functions are built from. No types, no factories.
- **`sandbox/lib/<object>/`** — one package per object the lib creates (factories plus own `New(...) api.<Object>` aggregate). None exist yet.
- **`examples/libraryExamples/<name>/<name>.go`** — outside the sandbox; self-contained `package main` programs calling `lib.New` directly, e.g. with `os.Args[1:]`.

### The Unused Mechanic

Every argument in `api.Lib.Args` starts out unread; `Used` tracks (index-for-index) which have been consumed. Any `Get*`/`IsPresent` call marks the argument(s) it matched as used — except `GetOptionsSize`/`GetKeyValuesSize`, which only count. `GetNextStringArg` (and typed variants) return the first still-unused argument in order, so a caller can read every expected flag/option first, then drain whatever positional arguments are left over. See `docs/References/UnusedMechanic.md`.

## Critical: this repo is documentation-driven

Changes are governed by required-reading docs, and several actions **must** update companion files in the same commit. Documentation is organized by **theme**, one index page per theme under `docs/Index/` (`LibUsage`, `Development`, `Templating`); the README links only to those indexes. Tutorials live flat in `docs/Tutorials/`, references and explanations flat in `docs/References/`.

| If you... | Read | And keep in sync |
|-----------|------|------------------|
| write or edit any `<Field>Factory` | `docs/References/Specs/Factories/Specs.md` | the `New` constructor in `sandbox/lib/new.go` that must call it |
| add/rename/delete any file or dir | `docs/References/Structure.md` | `docs/References/Structure.md` |
| add/rename/delete a `.md` file | `docs/Tutorials/HandleDocuments.md` | the theme index in `docs/Index/` listing it |
| add a lib function/object | `docs/Tutorials/HandleLibElements.md` | `docs/References/PublicApi.md` (+ detail page in `docs/References/PublicApi/`) |
| need an OS/third-party input inside `sandbox/` | `docs/References/SandboxIsolation.md` | pass it as a plain argument to `lib.New` or the relevant lib function, read from `examples/libraryExamples/` |
| add/rename/delete a sample | `docs/Tutorials/HandleLibrarySamples.md` | `docs/References/ApiSamplesList.md` |
| fork or adapt the template into a real library | `docs/Tutorials/ForkTemplate.md`, `AdaptExistingLib.md` | `docs/References/TemplateFileActions.md` (the per-file copy/create/rewrite/delete list both tutorials follow) |

`docs/References/Specs.md` is the index of every file specification; `AGENTS.md` points here and instructs: open the theme index matching your goal and check for a matching tutorial before acting, then locate the file's spec in `docs/References/Specs.md` (never browse `docs/References/Specs/` directly) before creating or editing it. A new public lib function or object must be declared in `sandbox/contracts/api/api.go`, given a factory in `sandbox/lib/`, **and** called from that package's `New` constructor, or callers get a nil field — that's the most common footgun, and `go build` will not catch it.

## Conventions

- Code that consumes the library from outside it (`examples/libraryExamples/`, third-party callers) aliases every import with the `verb` prefix: `verblib` (`sandbox`), `verbtypes` (`sandbox/contracts/api`). Files belonging to the library itself — `sandbox/` — keep the plain package names. See the `LibraryExamples` spec.
- Module path is `github.com/MateusMoutinhoOrg/Verb`; renaming it is a documented procedure — see `docs/Tutorials/RenameModule.md`.
- Public-facing lib API entries each get a detail page under `docs/References/PublicApi/` named `<pkg>.<Symbol>.md`.
- `docs/References/Specs/` holds the specifications: one directory per kind of file, each pairing a `Specs.md` (how the file must be shaped) with a `sample`. Never browse it — always locate a spec through `docs/References/Specs.md`.
