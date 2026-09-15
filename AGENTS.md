# AGENTS.md

## What this is

Verb is an **OS-independent argv parser**: flags, options, `key=value` arguments and
positionals read out of any `[]string`, with every argument that nothing asked for handed
back in order — the [Unused Mechanic](docs/UnusedMechanic/doc.md).

## This repository is generated

It is an [agnos](https://github.com/MateusMoutinhoOrg/Agnos) project. **agnos owns every
generated file**, and a hand edit to one is lost on the next build. Before changing
anything, read the pages that govern it — they are generated too, and they are the
authority over anything written here:

| Read | For |
|---|---|
| [docs/Rules/doc.md](docs/Rules/doc.md) | every rule, including the ones `agnos verify` enforces |
| [docs/Workflow/doc.md](docs/Workflow/doc.md) | the agnos command that makes each kind of change |
| [docs/Structure/doc.md](docs/Structure/doc.md) | what lives where |
| [docs/GeneratedFiles/doc.md](docs/GeneratedFiles/doc.md) | which files `build` rewrites |

```bash
agnos build        # verify + regenerate everything + go mod tidy + compile
agnos verify       # the schema check alone, writes nothing
agnos exec-test    # run every example and check it against its golden
```

Run `agnos build` after every hand edit. It is idempotent.

## What is particular to this repository

- **No `_test.go` anywhere under `sandbox/`.** The sandbox may import only sandbox
  packages, `testing` included, so `verify` refuses one. Behaviour is asserted by an
  example and its golden: add one with `agnos add-lib-example <name>`, then
  `agnos exec-test`. [`examples/lib/timestamps`](examples/lib/timestamps/example.go) is the
  regression check for the hand-written RFC 3339 parser.
- **An example is given its argv as a literal**, not from `os.Args[1:]` — it runs under
  `exec-test` with no arguments, and a literal reads as documentation besides. Each one
  writes what it parsed into `TestDir` and copies it into `AssertDir`; an example that
  copies nothing out fails.
- **No `time.Time` in `sandbox/api/`.** An instant crosses as an `int64` of nanoseconds, so
  the contract stays convertible for a consumer copying it — see
  [docs/Timestamps/doc.md](docs/Timestamps/doc.md) and
  [docs/EmbedVerb/doc.md](docs/EmbedVerb/doc.md). `sandbox/internal/rfc3339` implements the
  grammar in arithmetic, with no import but the api.
- **Adding a getter to `api.Parser` means adding its factory call to `NewParser`** in
  `sandbox/internal/argv/new.go`. A field no factory fills is a nil func that panics on
  first call, and nothing in the compiler or in `verify` catches it.
- **The four typed variants of a getter family share one reader.** `String`, `Int`,
  `Double` and `Timestamp` differ in the parse they run over the text and in nothing else,
  so a new family is one reader in `match.go` plus four factories.
