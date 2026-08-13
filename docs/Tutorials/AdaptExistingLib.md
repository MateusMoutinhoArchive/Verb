# Adapt a Pre-Existing Library

## Description
Covers converting a library that already exists into this project's struct-of-functions structure. To start a new library from scratch, follow [ForkTemplate.md](/docs/Tutorials/ForkTemplate.md) instead.

### Rules
- Read [Structure.md](/docs/References/Structure.md) and [Specs.md](/docs/References/Specs.md) before starting.
- Keep the separation defined in [Structure.md](/docs/References/Structure.md): public contract structs in `sandbox/contracts/`, the factories and their helpers in `sandbox/lib/`, and the entry point in `sandbox/`. Contracts are structs of function fields, never interfaces — see [StructContracts.md](/docs/References/StructContracts.md).
- Every file of the template has one action — **Copy**, **Create**, **Rewrite**, or **Delete**. Take it from [TemplateFileActions.md](/docs/References/TemplateFileActions.md); the steps below follow that order.
- The pre-existing package layout does **not** survive: all library logic ends up in `sandbox/lib/`. Code left in its original packages, or still calling `os`/`net`/third-party APIs directly, is not adapted — any such effect must arrive as a plain argument passed in from `examples/libraryExamples/`, never read inside `sandbox/`.
- Every public type the library returns becomes a contract struct in `sandbox/contracts/api`, whose function fields are filled by factories in `sandbox/lib/`. A type still declared in `sandbox/lib/` and handed back to callers is not adapted.
- Every file created or rewritten — code and `.md` alike — must follow its specification, located through [Specs.md](/docs/References/Specs.md).
- The adaptation is not complete until the final checklist in the last workflow step passes.

---

## Workflow
1. Recreate this project's directory layout inside the library being converted, using [Structure.md](/docs/References/Structure.md) as reference.
2. Copy every **[Copy](/docs/References/TemplateFileActions.md#copy)** file into the library unchanged — the specifications, tutorials, explanations, and [sandbox/new.go](../../sandbox/new.go).
3. Rewrite `sandbox/contracts/api/api.go` with the `Lib` struct and one struct per type the library hands back, following [HandleLibElements.md](/docs/Tutorials/HandleLibElements.md).
4. Rewrite the existing library code into `sandbox/lib/`: move each source file in, turn each public function into a `<Field>Factory(l *api.Lib)` that returns a closure for the matching api field, assign every factory's return value from the package's `New` constructor, and replace **every** OS-bound or third-party call with a plain argument passed into `lib.New`, following [HandleLibElements.md](/docs/Tutorials/HandleLibElements.md). Do not keep the code in its original packages, leave methods on internal types, or leave direct calls in place.
5. Create the samples in `examples/libraryExamples/` demonstrating the converted entry points, following [HandleLibrarySamples.md](/docs/Tutorials/HandleLibrarySamples.md).
6. Create the detail pages in `docs/References/PublicApi/` and rewrite `docs/References/PublicApi.md`, following [HandleLibElements.md](/docs/Tutorials/HandleLibElements.md#expose-in-the-public-api).
7. Delete every **[Delete](/docs/References/TemplateFileActions.md#delete)** file carried over from the template, plus the pre-existing code the converted lib replaced. For `.md` files, follow [HandleDocuments.md](/docs/Tutorials/HandleDocuments.md#delete-a-document).
8. Rewrite `docs/References/Structure.md` to describe the library's actual layout.
9. Create the tutorials specific to the converted library — one page per workflow its maintainers will repeat (e.g. adding a domain object, extending a feature, releasing) — following [HandleDocuments.md](/docs/Tutorials/HandleDocuments.md#add-a-document) and the [TutorialDocs specification](/docs/References/Specs/TutorialDocs/Specs.md). The template tutorials copied in step 2 cover the structure only; they do not document the library's own use cases.
10. Create any reference page the library needs beyond the public API — following [HandleDocuments.md](/docs/Tutorials/HandleDocuments.md#add-a-document) and the [ReferenceDocs specification](/docs/References/Specs/ReferenceDocs/Specs.md).
11. Rewrite the `README.md` and the theme indexes in `docs/Index/`: the README carries the overview, the badges, and the theme table; each index lists its theme's pages, and `docs/References/ApiSamplesList.md` lists the samples.
12. Verify the result:
```bash
go build ./...
```
Then confirm every item below — the adaptation is only done when all pass:
- All library logic lives in `sandbox/lib/`; no file there imports `os`, `net`, or a third-party implementation directly — every such input arrives as a plain argument to `lib.New` or a lib function, read from `examples/libraryExamples/`.
- `sandbox/contracts/api/api.go` declares every public object as a struct, and every one of its function fields is filled by a factory registered in that package's `New` constructor.
- `sandbox/new.go` is the only wiring point, and it imports nothing outside `sandbox/`.
- Tutorials and reference pages specific to this library exist under `docs/Tutorials/` and `docs/References/`.
- Every created or rewritten file matches its specification from [Specs.md](/docs/References/Specs.md).
- Every `.md` page is listed by a theme index in `docs/Index/`, the README links to every index, and `ApiSamplesList.md` lists every sample.
