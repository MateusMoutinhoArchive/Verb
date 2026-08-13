# Run API Examples

## Description
How to run the API examples provided in the repository to understand how to consume the project's Go API.

---

## Run API Examples

API examples are Go programs that demonstrate how to consume the library from code: each hands its process argv to `lib.New` and reads the fields of the returned `api.Lib`.

### Workflow

1. Browse the `/examples/libraryExamples/` directory for a package you want to explore (e.g., `Options`).
2. Run the package using `go run` from the project root, passing whatever arguments the sample parses:
   ```bash
   go run ./examples/libraryExamples/Options/Options.go --username alice
   ```
3. Examine the source code of the example to understand how the library API is invoked.
