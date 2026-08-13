# Verb

[![Go Reference](https://pkg.go.dev/badge/github.com/MateusMoutinhoOrg/Verb.svg)](https://pkg.go.dev/github.com/MateusMoutinhoOrg/Verb)
[![Release](https://img.shields.io/github/v/release/MateusMoutinhoOrg/Verb)](https://github.com/MateusMoutinhoOrg/Verb/releases/latest)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue)](go.mod)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

An OS-independent Go **argv parser** library demonstrating a **struct-of-functions** public API.

---

## Overview

Verb is a structured Go template that showcases how to build libraries exposed as plain data instead of interfaces. It uses a **struct-of-functions** pattern in which:

- **`/sandbox/contracts/`** declares the `api` structs the library hands back, whose fields are functions.
- **`/sandbox/lib/`** holds the pure library logic as factories filling those `api` structs — it never imports an OS-bound package.
- **`/sandbox/`** is the entry point: it takes the argument vector to parse and returns an `api.Lib`.
- **`/examples/libraryExamples/`** are runnable programs: each reads the real process argv and calls those fields.

This design ensures the library remains portable, testable, and easy to extend without modifying its core.

---

## Doc Index

Documentation is split into three themes, one index page each under `docs/Index/`, listing that theme's **Tutorials** — step-by-step workflows — and its **References** — explanations and lookups. Start from the theme index matching what you want to do.

| Theme | Description |
| --- | --- |
| [Library Usage](/docs/Index/LibUsage.md) | For library consumers: installing the module, parsing argv, and the Go API. |
| [Development](/docs/Index/Development.md) | For contributors: the mechanics, the workflows, and the specifications. |
| [Templating](/docs/Index/Templating.md) | For template users: forking, renaming, and adapting this structure into a new library. |

New here? [Library Usage → LibInitialization.md](/docs/Tutorials/LibInitialization.md) installs the module and runs a first program.

---

## License

This project is licensed under the [MIT License](./LICENSE).
