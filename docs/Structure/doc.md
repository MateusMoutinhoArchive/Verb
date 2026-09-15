# Structure

`(gen)` = written by `build`, never edited — the full list is in
[GeneratedFiles](../GeneratedFiles/doc.md).

```
adapters/  -->  sandbox/  <--  cmd/
(reaches OS)    (closed)       (wires)
```

Every line below is one entry of `AgnosConfig/structure.yaml` — add `<path>:
{description: "..."}` there, nested under `children:` of its parent, with `dir: true` on a
directory, `gen: true` on a file `build` rewrites, and `order:` to place it among its siblings
(unordered siblings follow, alphabetically).

```
AgnosConfig/      written once by `start`, read by every `build`
sandbox/          closed: imports nothing outside sandbox/, no OS packages
  new.go          (gen) New(deps) *api.Sandbox, one <x>.New<X> per api/ file
  api/            contracts only; imports nothing but sandbox/deps
    sandbox.go    (gen) one field per other file of api/, plus Deps
    argv.go       the Argv contract, plus the Parser one argument vector is bound to
    info.go       the Info contract - the library's own name and version
  deps/           one dir per capability the sandbox needs from outside; each imports nothing at all
    deps.go       (gen) one field per sub-directory, title-cased
    std/          formatting and the error constructor - `fmt` may not be named inside the sandbox
    stringsdeps/  prefix matching and number conversion - `strings` and `strconv` may not be named inside the sandbox
  internal/       the logic; unreachable from outside the sandbox
    config/       (gen) ProjectName and Version, from AgnosConfig/project.yaml
    argv/         NewArgv, the per-call NewParser, and the readers the four typed variants of every getter share
    rfc3339/      RFC 3339 text to nanoseconds since the epoch, in arithmetic alone - the sandbox may not name `time`
    info/         NewInfo
adapters/         the only place OS-bound and third-party code lives
  libs/           one dir per adapter, each exporting Bind and carrying an adapter.yaml
    std/          std over fmt, time, runtime and os
    stringsdeps/  stringsdeps over strings and strconv
  availables/     one dir per selection of adapters; new.go beside each is generated from its available.yaml
    standard/     std + stringsdeps - the whole of what Verb asks of the outside world
examples/         one dir per example under lib/, each a package main checked against its golden
  lib/            <name>/example.go + props.yaml, and the result.yaml exec-test writes
docs/             one dir per doc, holding doc.md + props.yaml. README.md indexes them all
AGENTS.md         what an agent reads first; CLAUDE.md is one line pointing here
go.mod            written by `start`; add-dep and remove-dep edit its require block
README.md         (gen) `render AgnosConfig/docs/ReadmeHeader.md` + the documentation index
```

Every rule this shape has to hold to — layers, naming, generated files, docs — is in
[Rules](../Rules/doc.md); the command that makes each change is in
[Workflow](../Workflow/doc.md).
