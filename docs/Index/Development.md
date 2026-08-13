# Development

## Description
Index of the documentation for contributors changing this repository: the mechanics every change runs into, the per-goal workflows, and the specifications every file must satisfy. Using the project is indexed by [LibUsage.md](/docs/Index/LibUsage.md); turning the project into a new library is indexed by [Templating.md](/docs/Index/Templating.md).

> [!IMPORTANT]
> **Read before contributing.** [Structure.md](/docs/References/Structure.md) and [Specs.md](/docs/References/Specs.md) are required reading: they say **where** a change belongs and **how** the file you touch must be shaped.

---

## Tutorials

- [HandleLibElements.md](/docs/Tutorials/HandleLibElements.md)
  - **description:** Add a function or object to the library: declare, write the factory, register, publish
  - [Add a Library Function](/docs/Tutorials/HandleLibElements.md#add-a-library-function)
  - [Add a Library Object](/docs/Tutorials/HandleLibElements.md#add-a-library-object)
  - [Expose in the Public API](/docs/Tutorials/HandleLibElements.md#expose-in-the-public-api)
- [HandleLibrarySamples.md](/docs/Tutorials/HandleLibrarySamples.md)
  - **description:** Create and run executable Go samples demonstrating a library feature
  - [Run a Library Sample](/docs/Tutorials/HandleLibrarySamples.md#run-a-library-sample)
  - [Add a Library Sample](/docs/Tutorials/HandleLibrarySamples.md#add-a-library-sample)
- [HandleDocuments.md](/docs/Tutorials/HandleDocuments.md)
  - **description:** Create, rename, move, or delete a `.md` file without leaving broken references
  - [Add a Document](/docs/Tutorials/HandleDocuments.md#add-a-document)
  - [Rename or Move a Document](/docs/Tutorials/HandleDocuments.md#rename-or-move-a-document)
  - [Delete a Document](/docs/Tutorials/HandleDocuments.md#delete-a-document)

---

## References

- [Structure.md](/docs/References/Structure.md)
  - **description:** The project's schema: which kind of file lives where, and its spec
  - [Root](/docs/References/Structure.md#root)
  - [`/sandbox/`](/docs/References/Structure.md#sandbox)
  - [`/examples/libraryExamples/`](/docs/References/Structure.md#exampleslibraryexamples)
  - [`/docs/`](/docs/References/Structure.md#docs)
- [Specs.md](/docs/References/Specs.md)
  - **description:** Index of every specification and the files each one governs
  - [Documentation Specifications](/docs/References/Specs.md#documentation-specifications)
  - [Code Specifications](/docs/References/Specs.md#code-specifications)
  - [Workflow](/docs/References/Specs.md#workflow)
- [Specs/](/docs/References/Specs/)
  - **description:** The specifications themselves — always reached through `Specs.md`, never browsed
- [SandboxIsolation.md](/docs/References/SandboxIsolation.md)
  - **description:** The sandbox wall: what `sandbox/` may not import, and why inputs are arguments
  - [The Two Trees](/docs/References/SandboxIsolation.md#the-two-trees)
  - [What the Wall Forbids](/docs/References/SandboxIsolation.md#what-the-wall-forbids)
  - [What the Wall Forbids in the Other Direction](/docs/References/SandboxIsolation.md#what-the-wall-forbids-in-the-other-direction)
  - [Why the Entry Point Lives Inside](/docs/References/SandboxIsolation.md#why-the-entry-point-lives-inside)
  - [If a Future Dependency Needs a Door](/docs/References/SandboxIsolation.md#if-a-future-dependency-needs-a-door)
- [StructContracts.md](/docs/References/StructContracts.md)
  - **description:** Why every contract is a struct of function fields, and how factories fill them
  - [The Shape](/docs/References/StructContracts.md#the-shape)
  - [Factories Fill the Fields](/docs/References/StructContracts.md#factories-fill-the-fields)
  - [Consuming a Library That Uses This Pattern](/docs/References/StructContracts.md#consuming-a-library-that-uses-this-pattern)
  - [What It Costs](/docs/References/StructContracts.md#what-it-costs)
- [UnusedMechanic.md](/docs/References/UnusedMechanic.md)
  - **description:** How every argument is marked used, and how leftovers are drained
  - [What Used Tracks](/docs/References/UnusedMechanic.md#what-used-tracks)
  - [Who Marks What](/docs/References/UnusedMechanic.md#who-marks-what)
  - [Draining the Leftovers](/docs/References/UnusedMechanic.md#draining-the-leftovers)
- [PublicApi.md](/docs/References/PublicApi.md)
  - **description:** Index of every public-facing entry of the library, grouped by kind
  - [Structs](/docs/References/PublicApi.md#structs)
  - [Functions](/docs/References/PublicApi.md#functions)
  - [Fields](/docs/References/PublicApi.md#fields)
