# Library Usage

## Description
Index of the documentation for developers consuming Verb as a Go library: handing the argument vector to the closed sandbox, parsing flags, options and positional arguments from code, and looking up the public API. Changing the library is indexed by [Development.md](/docs/Index/Development.md); turning it into a library of your own is indexed by [Templating.md](/docs/Index/Templating.md).

The library is always built the same way: the caller reads the real process argv, `lib.New` takes it into the closed sandbox, and the returned `api.Lib` carries every behavior as a function field.

---

## Tutorials

- [LibInitialization.md](/docs/Tutorials/LibInitialization.md)
  - **description:** Install the lib, call lib.New with argv, and run a first program
- [ParseOption.md](/docs/Tutorials/ParseOption.md)
  - **description:** Check a flag with IsPresent and read its value with GetStringOption
- [UseUnusedMechanic.md](/docs/Tutorials/UseUnusedMechanic.md)
  - **description:** Drain the leftover positional arguments with GetNextStringArg
- [RunApiSample.md](/docs/Tutorials/RunApiSample.md)
  - **description:** Run one of the shipped Go examples from the source tree
  - [Run API Examples](/docs/Tutorials/RunApiSample.md#run-api-examples)

---

## References

- [PublicApi.md](/docs/References/PublicApi.md)
  - **description:** Index of every public-facing entry of the library, grouped by kind
  - [Structs](/docs/References/PublicApi.md#structs)
  - [Functions](/docs/References/PublicApi.md#functions)
  - [Fields](/docs/References/PublicApi.md#fields)
- [ApiSamplesList.md](/docs/References/ApiSamplesList.md)
  - **description:** Every Go example shipped in `examples/libraryExamples/`
  - [Examples](/docs/References/ApiSamplesList.md#examples)
- [UnusedMechanic.md](/docs/References/UnusedMechanic.md)
  - **description:** How every argument is marked used, and how leftovers are drained
  - [What Used Tracks](/docs/References/UnusedMechanic.md#what-used-tracks)
  - [Who Marks What](/docs/References/UnusedMechanic.md#who-marks-what)
  - [Draining the Leftovers](/docs/References/UnusedMechanic.md#draining-the-leftovers)
- [StructContracts.md](/docs/References/StructContracts.md)
  - **description:** Why every contract is a struct of function fields, and how factories fill them
  - [The Shape](/docs/References/StructContracts.md#the-shape)
  - [Factories Fill the Fields](/docs/References/StructContracts.md#factories-fill-the-fields)
  - [Consuming a Library That Uses This Pattern](/docs/References/StructContracts.md#consuming-a-library-that-uses-this-pattern)
  - [What It Costs](/docs/References/StructContracts.md#what-it-costs)
