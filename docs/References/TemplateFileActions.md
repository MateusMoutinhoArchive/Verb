# Template File Actions

## Description
Lists every file and directory of this template and the action it takes when the template is forked into a new library or an existing library is adapted to it. Each file falls into exactly one action: **Copy**, **Create**, **Rewrite**, or **Delete**. The workflows using this list are [ForkTemplate.md](/docs/Tutorials/ForkTemplate.md) and [AdaptExistingLib.md](/docs/Tutorials/AdaptExistingLib.md).

---

## Copy

Taken as-is from the template. They describe the structure itself, not the library, so they carry over unchanged. Adapting them is allowed but never required.

Copying these files carries over the template's **generic** guides and specifications only. The new library must still **create** its own case-specific tutorials and reference pages — see [Create](#create).

| Path | Description |
|------|-------------|
| `docs/References/Specs/*` | The specifications every file of the new library must be shaped by |
| `docs/References/Specs.md` | The index locating each specification |
| `docs/References/TemplateFileActions.md` | This page |
| `docs/Tutorials/*` | The workflow guides |
| `docs/References/SandboxIsolation.md`, `StructContracts.md` | The explanations of the structure's mechanics |
| `docs/Index/*` | The theme indexes, updated as the new library's pages replace the template's |
| `sandbox/new.go` | The `New` constructor delegating to the `sandbox/lib` constructor |

---

## Create

Written from scratch for the library being built or adapted. Nothing of the template's content survives here — the example files occupying these paths are removed by **[Delete](#delete)**. Every created file must be shaped by the specification in its row.

| Path | Description | Specification |
|------|-------------|---------------|
| `sandbox/lib/new.go`, `sandbox/lib/publicfunctions/*` | The lib's field factories and the `New` constructor running them all | [LibFunctions](/docs/References/Specs/LibFunctions/Specs.md) |
| `sandbox/lib/<object>/*` | One package per object the library hands back: its field factories and the `New` constructor running them all | [LibObjects](/docs/References/Specs/LibObjects/Specs.md) |
| `docs/References/PublicApi/*` | One detail page per public API entry | [ReferenceDocs](/docs/References/Specs/ReferenceDocs/Specs.md) |
| `docs/References/<Name>.md` | Any reference page the new library needs beyond the public API index | [ReferenceDocs](/docs/References/Specs/ReferenceDocs/Specs.md) |
| `docs/Tutorials/<Goal>.md` | One tutorial per workflow specific to the new library — the template tutorials carried over by **[Copy](#copy)** do **not** fulfill this | [TutorialDocs](/docs/References/Specs/TutorialDocs/Specs.md) |
| `examples/libraryExamples/<example>/<example>.go` | One runnable sample per demonstrated use case | [LibraryExamples](/docs/References/Specs/LibraryExamples/Specs.md) |
| `sandbox/config/*` | The compile-time constants the new library reads | |

---

## Rewrite

Kept in place, with their content replaced by the new library's. The file keeps its path and its shape; only what it declares or documents changes. Every rewritten file must be shaped by the specification in its row.

| Path | Rewrite with | Specification |
|------|--------------|---------------|
| `README.md` | The new library's overview, badges, and theme table | [Readme](/docs/References/Specs/Readme/Specs.md) |
| `sandbox/contracts/api/api.go` | The `Lib` struct and one struct per object the new library hands back | [Outputs](/docs/References/Specs/Outputs/Specs.md) |
| `docs/References/PublicApi.md` | The index of the new public API entries | [ReferenceDocs](/docs/References/Specs/ReferenceDocs/Specs.md) |
| `docs/References/Structure.md` | The layout of the new library | [Structure](/docs/References/Specs/Structure/Specs.md) |
| `docs/References/ApiSamplesList.md` | The list of the new library's samples | [ReferenceDocs](/docs/References/Specs/ReferenceDocs/Specs.md) |
| `docs/Index/<Theme>.md` | Each theme's entry list, once the new library's pages exist | [Index](/docs/References/Specs/Index/Specs.md) |

---

## Delete

The template's example content. Removed once the new library's own files exist.

| Path | Description |
|------|-------------|
| `sandbox/lib/*` | The example lib factories, helper packages, and object packages |
| `docs/References/PublicApi/*` | The example API detail pages |
| `examples/libraryExamples/*` | The example samples |
