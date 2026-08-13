# Project Structure

This document maps the project **schema** — the kinds of files the project is built from — not every concrete file. A slot with a **Spec** name is governed by a specification; resolve the name through [Specs.md](/docs/References/Specs.md) to get its description and sample.

The project is a **library** and nothing else: it ships no binary and no interface of its own. It is split into two top-level code trees, and the dependency flow between them is one-way:

```
sandbox/  ◀──  examples/libraryExamples/
(closed)       (consume the lib)
```

- **`/sandbox/`** is a **closed sandbox**: the pure library. Nothing inside it may import `examples/libraryExamples/`, a third-party module, or any OS-bound standard-library package — every input it needs arrives as a plain function argument. See [SandboxIsolation.md](/docs/References/SandboxIsolation.md).
- **`/examples/libraryExamples/`** sits outside the sandbox and is the only place `os.Args` is read and handed to the library — the same wiring a consuming project writes for itself.

Because every behavior is a field of `api.Lib`, a consumer never calls into the sandbox's internals: it hands the argument vector to `lib.New` and reads the struct it gets back.

## Root

| File | Description | Spec |
|------|-------------|------|
| `README.md` | Project overview and the Doc Index pointing at each theme index under `docs/Index/` | Readme |
| `LICENSE` | License terms for the project | |
| `go.mod` | Go module definition and dependencies | |
| `.gitignore` | Intentionally untracked files to ignore | |

---

## `/sandbox/`
The closed sandbox — the pure library. It holds its own entry point, the contracts everything is wired through, the configuration constants, and the implementation. It reaches nothing outside itself. Its package is named `lib`, so consumers import it as `lib "…/sandbox"` and call `lib.New`.

| File | Description | Spec |
|------|-------------|------|
| `new.go` | The `New(args []string) api.Lib` constructor delegating to the `sandbox/lib` constructor | |

### `/sandbox/contracts/`
The structs the rest of the project is wired through — the only part of the sandbox anything outside it may import. Contracts hold the project's **public types** and are structs of function fields, never interfaces; see [StructContracts.md](/docs/References/StructContracts.md). Contracts import nothing from `sandbox/lib/`.

#### `/sandbox/contracts/api/`
The structs the library hands back to callers.

| File | Description | Spec |
|------|-------------|------|
| `api.go` | The `Lib` entry-point struct plus one struct per object the lib creates | Outputs |

### `/sandbox/config/`
Static configuration the sandbox reads at compile time: the fixed text and numbers the library parses or reports. Holding them as Go constants rather than as files keeps every reference under the compiler's eye — a renamed constant is a build failure rather than a blank line at runtime — and costs no read at all. Nothing outside the sandbox imports this package.

| File | Description | Spec |
|------|-------------|------|
| `timestamp.go` | The `TimestampLayout` constant every Timestamp getter parses its value with | |

### `/sandbox/lib/`
The entry-point implementation and every internal package the library is built from. Each package here holds the functions that take a pointer to an [`api`](#sandboxcontractsapi) struct and return closures reading that struct's own state, which the package's constructor assigns into the matching function fields. Types never live here; they stay in `contracts/`.

| File | Description | Spec |
|------|-------------|------|
| `new.go` | The `New(args []string) api.Lib` constructor that assigns every factory's return value and runs them all | LibFunctions |

#### `/sandbox/lib/publicfunctions/`
One file per public function field of `api.Lib`. Each file holds its `<Field>Factory(l *api.Lib)` that returns a closure. The `New` constructor in `sandbox/lib/new.go` calls every factory here to fill `api.Lib`.

| File | Description | Spec |
|------|-------------|------|
| `<Function>.go` | One file per lib function, holding its `<Field>Factory(l *api.Lib)` that returns a closure | LibFunctions |

#### `/sandbox/lib/<object>/`
One package per object the library creates, named after the object itself. None exist yet in this template.

| File | Description | Spec |
|------|-------------|------|
| `<object>.go` | The object's `<Field>Factory` functions, each returning a closure, plus the `New(…) api.<Object>` constructor that assigns every factory's return value | LibObjects |

#### `/sandbox/lib/argv/`
Shared helpers over the argument vector, used by every package under `sandbox/lib/`: matching a flag or a `key=` prefix, locating an occurrence, marking it used, and parsing the raw text into a typed value. It declares **no types and no factories** — it is the one internal package that is neither an object nor the entry point, so no specification governs it.

| File | Description | Spec |
|------|-------------|------|
| `argv.go` | The matching, consuming, and parsing helpers the public functions are built from | |

---

## `/examples/libraryExamples/`
Outside the sandbox. Runnable Go examples demonstrating how to use the library from code — the only place `os.Args` is read.

### `/examples/libraryExamples/<example>/`

| File | Description | Spec |
|------|-------------|------|
| `<example>.go` | Self-contained `package main` calling `lib.New` with the real process argv | LibraryExamples |

**Run an example:**
```sh
go run ./examples/libraryExamples/<example>/<example>.go
```

---

## `/docs/`
Documentation of the project, split by **kind of page**: `Index/` holds one entry point per theme, `Tutorials/` holds every workflow, `References/` holds every lookup and explanation. A **theme** — what the reader wants to accomplish — is not a directory: it is the index that lists a page. `Tutorials/` and `References/` are flat, so a page's file name must be unique inside the directory it lands in. The [README](/README.md) links to the three indexes and to nothing else inside `docs/`.

| Directory | Description |
|-----------|-------------|
| `Index/` | One entry point per theme, each listing the pages of its theme |
| `Tutorials/` | Every workflow page of the project, whatever theme it belongs to |
| `References/` | Every lookup and explanation page, whatever theme it belongs to |

### `/docs/Index/`
One page per theme. The three themes are `LibUsage` — installing the module, parsing argv from Go, and the public API; `Development` — contributing: the mechanics, the workflows, the specifications; and `Templating` — turning this repository into another library.

| File | Description | Spec |
|------|-------------|------|
| `<Theme>.md` | The theme's entry point: its Tutorials and its References, each entry listing that page's sections | Index |

### `/docs/Tutorials/`
One page per workflow, its title phrased as the action it performs. A page can belong to one or more themes, and the theme indexes in [`/docs/Index/`](#docsindex) are what say which.

| File | Description | Spec |
|------|-------------|------|
| `<Goal>.md` | One page per workflow, written as numbered steps a reader follows to the end | TutorialDocs |

### `/docs/References/`
One page per lookup table or explained mechanic, plus the two directories the project's biggest listings live in.

| File | Description | Spec |
|------|-------------|------|
| `<Name>.md` | One page per lookup table or explained mechanic | ReferenceDocs / ExplanationDocs |
| `PublicApi.md` | Index of all public-facing components, linking to their detail pages | ReferenceDocs |
| `ApiSamplesList.md` | Every example under `examples/libraryExamples/` | ReferenceDocs |
| `Structure.md` | The project's schema and the purpose of each component | Structure |
| `Specs.md` | Index of every specification and the files each one governs | |
| `SandboxIsolation.md` | What the sandbox may not import, and why every input is an argument | ExplanationDocs |
| `StructContracts.md` | Why contracts are structs of function fields, and how factories fill them | ExplanationDocs |
| `UnusedMechanic.md` | How arguments are marked used, and how leftovers are drained in order | ExplanationDocs |
| `TemplateFileActions.md` | The action each template file takes when forking or adapting | ReferenceDocs |

#### `/docs/References/PublicApi/`
One detail page per public-facing component, indexed by `PublicApi.md`. Reach a page through that index rather than by browsing the directory.

| File | Description | Spec |
|------|-------------|------|
| `<pkg>.<Symbol>.md` | One detail page per public struct, function, or field, named after the package the symbol is declared in | ReferenceDocs |

#### `/docs/References/Specs/`
The specifications describing how each kind of file in the project must be shaped. Never browse this directory — locate a specification by reading `Specs.md`.

| File | Description | Spec |
|------|-------------|------|
| `<Spec>/Specs.md` | The required shape of the artifact the specification governs | |
| `<Spec>/sample.<ext>` | Concrete reference implementation of the specification | |
