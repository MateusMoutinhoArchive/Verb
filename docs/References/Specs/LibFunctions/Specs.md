# LibFunctions Specification

## Description
Defines the required shape of a library function — a **factory** in `sandbox/lib/publicfunctions/` that returns the closure for one function field of the `Lib` struct declared in `sandbox/contracts/api/api.go`. A factory takes a pointer to the api struct and returns a closure; the caller — the `New` constructor in `sandbox/lib/new.go` — assigns it into the field, and the closure reaches the struct's own state through that pointer.

### Rules
- One **file** per function field, named `<Field>.go` after the field it fills, holding that field's factory and nothing else.
- One factory per function field, named `<Field>Factory` and taking a single `*api.Lib` parameter: `func IsPresentFactory(l *api.Lib) func(...)`. Its only job is to build and return the closure.
- The factory's body returns a closure for the field: `return func(...) ... { … }`. The closure's signature must match the field's declaration in `sandbox/contracts/api/api.go` exactly.
- Every factory must be called from `New(args []string) api.Lib` in `sandbox/lib/new.go`, which assigns its return value into the matching field and doubles as the factory aggregate — there is no separate `Factory` function. A field whose factory's return value is never assigned stays nil and panics on first call; the compiler does not check this.
- State is read **only** through the carrier pointer inside the closure — `l.<Field>(...)` — never captured at factory time. Reading `l.<field>` inside the closure rather than capturing it at factory time is what lets the struct stay authoritative.
- `sandbox/` is a closed sandbox: a factory must never import `examples/`, a third-party module, or an OS-bound standard-library package (`os`, `net`, `syscall`, …). See [SandboxIsolation.md](/docs/References/SandboxIsolation.md).
- A closure returning a library object returns that object's **api struct**, built by the object package's `New` constructor — see the [LibObjects](/docs/References/Specs/LibObjects/Specs.md) specification.
- A closure returning an optional object returns the api struct's **zero value** on the miss path; there is no nil struct to return.
- Logic shared by several factories — matching an argument, parsing a value — lives in a helper package under `sandbox/lib/` (e.g. `sandbox/lib/argv/`), never duplicated across factory files.
- Factories and the fields they fill must have doc comments, and the fields must be listed in [PublicApi.md](/docs/References/PublicApi.md).

## Structure
1. **Package clause**: `package publicfunctions`, in `sandbox/lib/publicfunctions/<Field>.go`.
2. **Field factory**: `func <Field>Factory(l *api.Lib) <FieldType>` returning a closure for `l.<Field>`, reading the struct's own state through `l`.
3. **Doc comment**: one sentence naming the field the factory fills and what the closure does.
4. **`New` constructor** (in `sandbox/lib/new.go`, `package lib`): `func New(args []string) api.Lib` building `api.Lib{...}`, calling every field factory of `publicfunctions` exactly once and assigning its return value into the matching field, and returning the struct. `sandbox/new.go` does nothing but delegate to it.

> **Note**: For a concrete example, refer to [sample.go](./sample.go).
