# LibObjects Specification

## Description
Defines the required shape of an object created by the library — a package under `sandbox/lib/<object>/` holding the **factories** that fill a struct declared in `sandbox/contracts/api/api.go`, plus the `New` constructor that builds it. The object is the api struct itself; there is no separate internal mirror type.

### Rules
- Each object lives in its own package under `sandbox/lib/`, named `<object>` after the object it fills. No `internal_` prefix — everything under `sandbox/lib/` is the library's own implementation, closed by convention.
- `sandbox/lib/` holds **only factories, constructors, and their helpers**. The object's type is declared in `sandbox/contracts/api`; a package under `sandbox/lib/` may import contracts, never the reverse.
- The object's api struct declares its data fields plus its function fields — see the [Outputs](/docs/References/Specs/Outputs/Specs.md) specification.
- One factory per function field, named `<Field>Factory` and taking a single `*api.<Object>` parameter, whose body returns a closure for that field. The closure reads the object's properties through the pointer, so they are current at call time.
- The package must expose a `New(...) api.<Object>` constructor that fills the object's data fields, calls **every** field factory and assigns its return value into the matching field, and returns the struct by value. `New` is the factory aggregate — there is no separate `Factory` function. A field no factory's return value is assigned into stays nil and panics on first call.
- The object is created only through a factory on `Lib` (e.g. `NewExampleObjectFactory`) whose returned closure calls `<object>.New(...)` — callers never build it directly.
- `sandbox/` is a closed sandbox: an object package must never import `examples/`, a third-party module, or an OS-bound standard-library package (`os`, `net`, `syscall`, …). See [SandboxIsolation.md](/docs/References/SandboxIsolation.md).
- Exported api structs, constructors, and factories must have doc comments, and the public fields must be listed in [PublicApi.md](/docs/References/PublicApi.md).

## Structure
1. **Struct declaration** (in `sandbox/contracts/api/api.go`): the object's data fields and the function fields it exposes — see [Outputs](/docs/References/Specs/Outputs/Specs.md).
2. **Field factories** (in `sandbox/lib/<object>/<object>.go`): `func <Field>Factory(o *api.<Object>) <FieldType>`, each returning one closure.
3. **`New` constructor**: `func New(...) api.<Object>` filling the data fields, calling every field factory exactly once and assigning its return value into the matching field, and returning the struct.
4. **Constructor factory** (in `sandbox/lib/publicfunctions/`): `func New<Object>Factory(l *api.Lib) func(...) api.<Object>` whose returned closure calls `<object>.New(...)`.

> **Note**: For a concrete example, refer to [sample.go](./sample.go).
