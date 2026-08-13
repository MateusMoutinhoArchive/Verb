# Verb

[![Go Reference](https://pkg.go.dev/badge/github.com/MateusMoutinhoOrg/Verb.svg)](https://pkg.go.dev/github.com/MateusMoutinhoOrg/Verb)
[![Release](https://img.shields.io/github/v/release/MateusMoutinhoOrg/Verb)](https://github.com/MateusMoutinhoOrg/Verb/releases/latest)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue)](go.mod)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

An OS-independent Go **library template** — an argv parser whose entire behavior lives inside a closed sandbox and is handed back as a struct of functions.

---

## Overview

Verb is a **full library template** designed to be completely independent of the underlying operating system. It provides a complete harness and architectural foundation for building Go libraries whose behavior is fully decoupled from the hosting environment, exposed as plain data instead of interfaces. Furthermore, the repository is designed to be self-teaching, providing **comprehensive tutorials for every kind of usecase** directly within its own documentation.

The core of the library lives in **`/sandbox/`**: a **closed sandbox** that reaches nothing outside itself — no third-party module, no OS-bound standard-library package. Everything it needs from the outside world arrives as a plain function argument.

```
sandbox/  ◀──  examples/libraryExamples/
(closed)       (consume the lib)
```

Nothing is exported but the library itself. There is no binary, no command, and no output of its own: a consumer hands the argument vector to `lib.New` and calls the fields of the `api.Lib` it gets back.

- **`/sandbox/`**: The closed library taking an `[]string` and returning an `api.Lib`.
- **`/examples/libraryExamples/`**: Places where the real process argv and the library are wired together.

The parser it demonstrates is deliberately small — flag presence, repeatable options, `key=value` pairs, positional arguments by index, and the leftovers drained in order — but every one of those is a function field a derived library replaces with its own. See [PublicApi.md](/docs/References/PublicApi.md) for every field.

See [SandboxIsolation.md](/docs/References/SandboxIsolation.md) and [StructContracts.md](/docs/References/StructContracts.md) for the full mechanic.

---

## Doc Index

Documentation is split into three themes, one index page each under `docs/Index/`, listing that theme's **Tutorials** — step-by-step workflows — and its **References** — explanations and lookups. Start from the theme index matching what you want to do.

| Theme | Description |
| --- | --- |
| [Library Usage](/docs/Index/LibUsage.md) | For library consumers: installing the module, parsing argv, and calling the Go API. |
| [Development](/docs/Index/Development.md) | For contributors: the mechanics, the per-goal workflows, and the specifications. |
| [Templating](/docs/Index/Templating.md) | For template users: forking, renaming, and adapting this structure into a new library. |

New here? [Library Usage → LibInitialization.md](/docs/Tutorials/LibInitialization.md) installs the module and runs a first program.

---

## License

This project is licensed under the [MIT License](./LICENSE).
