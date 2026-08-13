# Project Structure

This document maps the project **schema** — the kinds of files a project is built from — not every concrete file. A slot with a **Spec** name is governed by a specification; resolve the name through [Specs.md](/docs/References/Specs.md) to get its description and sample. Use them to build a project from scratch, even before any code exists.

## Root

| File | Description | Spec |
|------|-------------|------|
| `README.md` | Project overview and quick-start guide | Readme |
| `LICENSE` | License terms for the project | |
| `go.mod` | Go module definition | |

---

## `/docs/`
Documentation of the project.

### `/docs/References/`

| File | Description | Spec |
|------|-------------|------|
| `Structure.md` | The project's schema and the purpose of each component | Structure |
| `Specs.md` | Index of every specification and the files each one governs | |

#### `/docs/References/Specs/`
The specifications describing how each kind of file in the project must be shaped.

| File | Description | Spec |
|------|-------------|------|
| `<Spec>/` | One directory per specification, holding its `Specs.md` and `sample` | |

---

## `/sandbox/`
The closed sandbox — the library, its public entry point, its contracts, and its internal logic. It imports nothing outside itself.

| File | Description | Spec |
|------|-------------|------|
| `new.go` | The `New` constructor taking plain arguments and returning an `api.Lib` | |

### `/sandbox/contracts/`
The public contract structs the project is wired through — the only part of the sandbox the outside imports.

| File | Description | Spec |
|------|-------------|------|
| `api/api.go` | The structs the library hands back to callers | Outputs |

### `/sandbox/lib/`
The factories filling the `api` structs' function fields, plus the constructor aggregating them. Declares no types.

| File | Description | Spec |
|------|-------------|------|
| `new.go` | The `New` constructor assigning every factory's return value into `api.Lib` | LibFunctions |
| `publicfunctions/<Function>.go` | One file per public function field, holding its factory | LibFunctions |
| `<object>/<object>.go` | The factories filling an `api` struct's fields, plus its `New` constructor | LibObjects |

---

## `/examples/libraryExamples/`
Outside the sandbox. Runnable examples demonstrating how to use the library.

| File | Description | Spec |
|------|-------------|------|
| `<example>/<example>.go` | Self-contained `package main` calling `lib.New` | LibraryExamples |
