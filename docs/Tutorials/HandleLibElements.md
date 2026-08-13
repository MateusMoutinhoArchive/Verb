# Handle Lib Elements

## Description
Covers adding an element to the library and publishing it: a function field of `api.Lib`, an object the library creates, and the public API entry either one is documented through. Running the result is [HandleLibrarySamples.md](/docs/Tutorials/HandleLibrarySamples.md).

### Rules
- Every type in the project is declared in [sandbox/contracts/api/api.go](/sandbox/contracts/api/api.go). [sandbox/lib/](/sandbox/lib/) holds only factories, constructors, and their helpers.
- An element is only usable once its factory's return value is **assigned** from the `New` constructor that aggregates its package — an unassigned field stays nil and panics on first call, and the compiler does not catch it.
- One factory per field, named `<Field>Factory`, taking a single pointer to the carrier and returning one closure.
- State is read as `l.<field>` **inside** the closure, never captured at factory time — that is what keeps the struct authoritative.
- `sandbox/` is a closed sandbox: library code must never import [examples/](/examples/), a third-party module, or an OS-bound standard-library package (`os`, `net`, `syscall`, …). See [SandboxIsolation.md](/docs/References/SandboxIsolation.md).
- Adding a file or directory under [sandbox/lib/](/sandbox/lib/) requires updating [Structure.md](/docs/References/Structure.md).
- The file you write must follow its specification — locate it in [Specs.md](/docs/References/Specs.md).

---

## Add a Library Function

### Workflow
1. Declare the function as a field of the `Lib` struct in [sandbox/contracts/api/api.go](/sandbox/contracts/api/api.go):
   ```go
   type Lib struct {
       Args           []string
       Used           []bool
       IsPresent      func(flags []string) bool
       GetOptionsSize func(flags []string) int
       HasOption      func(flags []string) bool // new function
   }
   ```
2. Write its factory in a file of its own under [sandbox/lib/publicfunctions/](/sandbox/lib/publicfunctions/), named after the field (`HasOption.go`), with the identical signature, returning the closure:
   ```go
   package publicfunctions

   // HasOptionFactory fills api.Lib.HasOption with a closure reporting
   // whether an option occurs at least once, without consuming it.
   func HasOptionFactory(l *api.Lib) func(flags []string) bool {
       return func(flags []string) bool {
           return l.GetOptionsSize(flags) > 0
       }
   }
   ```
   > Calling another field from inside a closure (`l.GetOptionsSize` above) is fine: by the time `HasOption` runs, `New` has already filled every field.
   >
   > Logic shared with other factories — matching an argument, parsing a value — belongs in [sandbox/lib/argv/](/sandbox/lib/argv/), not duplicated per file.
3. Assign the factory's return value in `New`, in [sandbox/lib/new.go](/sandbox/lib/new.go) — without this line the field stays nil and the function panics when called:
   ```go
   func New(args []string) api.Lib {
       l := api.Lib{Args: args, Used: make([]bool, len(args))}
       l.IsPresent = publicfunctions.IsPresentFactory(&l)
       l.GetOptionsSize = publicfunctions.GetOptionsSizeFactory(&l)
       l.HasOption = publicfunctions.HasOptionFactory(&l) // register it
       return l
   }
   ```
4. If the function returns a new object, create it following [Add a Library Object](#add-a-library-object) and return the object's `api` struct.
5. Expose the function following [Expose in the Public API](#expose-in-the-public-api).
6. Register the new file in [Structure.md](/docs/References/Structure.md).
7. If the function needs a runnable demonstration, add one following [HandleLibrarySamples.md](/docs/Tutorials/HandleLibrarySamples.md).
8. Build the project — a signature mismatch fails here, but a forgotten assignment in `New` does not:
   ```bash
   go build ./...
   ```
9. Call the new field once — from a sample or a test — to confirm it is not nil. That is the only check that catches a missing assignment.

---

## Add a Library Object

### Workflow
1. Declare the object's struct in [sandbox/contracts/api/api.go](/sandbox/contracts/api/api.go). This walkthrough adds a `Bucket`, a namespaced view over a set of keys:
   ```go
   type Bucket struct {
       Prefix  string
       FullKey func(name string) string
   }
   ```
2. Create the object's package and file, both named after the object (e.g. `sandbox/lib/bucket/bucket.go`), holding its field factories and its `New` constructor:
   ```go
   package bucket

   import (
       "github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
   )

   // FullKeyFactory fills api.Bucket.FullKey with a closure namespacing a
   // name under the bucket's prefix.
   func FullKeyFactory(b *api.Bucket) func(name string) string {
       return func(name string) string {
           return b.Prefix + ":" + name
       }
   }

   // New builds an api.Bucket and runs every bucket factory over it,
   // assigning each return value into its matching function field.
   func New(prefix string) api.Bucket {
       b := api.Bucket{
           Prefix: prefix,
       }
       b.FullKey = FullKeyFactory(&b)
       return b
   }
   ```
3. Declare the constructor as a field of the `Lib` api struct in [sandbox/contracts/api/api.go](/sandbox/contracts/api/api.go), returning the object's api struct:
   ```go
   type Lib struct {
       NewBucket func(prefix string) Bucket
   }
   ```
4. Write the constructor's factory in [sandbox/lib/publicfunctions/](/sandbox/lib/publicfunctions/), in `NewBucket.go`:
   ```go
   // NewBucketFactory fills api.Lib.NewBucket with a closure creating a
   // Bucket.
   func NewBucketFactory(l *api.Lib) func(prefix string) api.Bucket {
       return func(prefix string) api.Bucket {
           return bucket.New(prefix)
       }
   }
   ```
5. Assign `NewBucketFactory`'s return value in `New`, exactly as step 3 of [Add a Library Function](#add-a-library-function) does.
6. Expose the object, its constructor, and its fields following [Expose in the Public API](#expose-in-the-public-api).
7. Register the new directory and file in [Structure.md](/docs/References/Structure.md).
8. If the object needs a runnable demonstration, add one following [HandleLibrarySamples.md](/docs/Tutorials/HandleLibrarySamples.md).
9. Build the project, then call the new field once to confirm it is not nil — a missing assignment in `New` compiles cleanly:
   ```bash
   go build ./...
   ```

---

## Expose in the Public API

### Rules
- Every public-facing entry must be listed in [PublicApi.md](/docs/References/PublicApi.md).
- Detail pages live in [docs/References/PublicApi/](/docs/References/PublicApi/) and are named `<pkg>.<Symbol>.md`.

### Workflow
1. Open [PublicApi.md](/docs/References/PublicApi.md).
2. Add the struct, function, or field to the section matching its kind, with a one-line description. An object is public only through its `sandbox/contracts/api` struct — never document a `sandbox/lib/` symbol as the entry.
3. Create the detail page under [docs/References/PublicApi/](/docs/References/PublicApi/), named `<pkg>.<Symbol>.md` after the package the symbol is declared in (e.g. `api.IsPresent.md`), following [HandleDocuments.md](/docs/Tutorials/HandleDocuments.md).
4. Link the new detail page from its entry in [PublicApi.md](/docs/References/PublicApi.md).
5. Register the detail page in [Structure.md](/docs/References/Structure.md) if it is a new structural component.
