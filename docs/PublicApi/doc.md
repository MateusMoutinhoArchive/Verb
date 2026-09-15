# PublicApi

Every exported symbol of `github.com/MateusMoutinhoOrg/Verb`, read straight from the contract sources on
every build: `sandbox/api/` is the surface `sandbox.New` returns, `sandbox/deps/`
the contracts an adapter fills and a caller may replace. Each description below is the
doc comment of the declaration itself — change the comment, run `build`, and this page
follows.

## Entry points

| Symbol | Signature |
| --- | --- |
| `sandbox.New` | `func(deps *deps.Deps) *api.Sandbox` |
| `standard.New` | `func() deps.Deps` (`adapters/availables/standard`) |

Implementations live under `sandbox/internal` and are unreachable: every contract is a
struct of function fields, filled by a binder.

# The sandbox api

## `sandbox/api/sandbox.go`

### `Sandbox`

Sandbox is the whole library: one field per contract declared in sandbox/api/, each built by the New<Contract> of its own package under sandbox/internal/. sandbox.New returns it, and nothing callable lives outside of it.

| Field | Type | Description |
| --- | --- | --- |
| `Deps` | `*deps.Deps` | Deps is every capability the sandbox reaches the outside world through. It rides on the api so that a function handed the Sandbox holds the whole of what it needs, and can call another field of the api besides — which is what makes a field a caller replaced take effect everywhere. It is also the one field that does not cross into a consumer: an installed copy of this contract carries the api, never the wiring behind it. |
| `Argv` | `Argv` |  |
| `Info` | `Info` |  |

## `sandbox/api/argv.go`

### `Argv`

Argv is the contract that turns an argument vector into a Parser, carried by the Sandbox as the field of the same name. A parser is bound to one argument vector and carries the read state of it, so it is built per call rather than once at wiring time — which is why this contract holds a constructor and nothing else.

| Field | Type | Description |
| --- | --- | --- |
| `New` | `func(args []string) Parser` | New builds a Parser bound to the given arguments, e.g. os.Args[1:]. The slice is kept as it stands and is never written to; every argument starts out unread. Passing nil or an empty slice is valid and yields a parser that matches nothing. |

### `Parser`

Parser is an argument-vector (argv) parser: the argv it was built over, the read state of it, and one function field per way of reading a value out of it. Every argument starts unread and every reader marks what it matched, so whatever is left over is exactly the positional arguments nothing asked for — docs/UnusedMechanic holds who marks what, and in what order to read. Each reader family (Option, Arg, NextArg, KeyValues) comes in four types: String (the raw text), Int (base-10), Double (a float64) and Timestamp (RFC 3339 text reported as nanoseconds since the Unix epoch, UTC — docs/Timestamps). A Get* function returns an error rather than a bool or a panic when it cannot produce a value; the two *Size functions and IsPresent return plain values, because counting and checking presence cannot fail.

| Field | Type | Description |
| --- | --- | --- |
| `Args` | `[]string` | Args is the argument vector being parsed, the same slice Argv.New was called with. Every index-based function refers to positions in this slice. Treat it as read-only: mutating it after construction leaves Used out of sync and produces undefined matching behaviour. |
| `Used` | `[]bool` | Used tracks, index for index against Args, which arguments have already been matched by a previous call. Used[i] is true once Args[i] has been consumed by any Get* function or by IsPresent. It starts entirely false and only grows more true over the parser's lifetime — the Unused Mechanic reads it to find the next never-consumed positional argument. Treat it as read-only. |
| `IsPresent` | `func(flags []string) bool` | IsPresent reports whether any of the given flag spellings (e.g. []string{"-q", "--quiet"}) occurs anywhere in the unread portion of Args. On a match it marks that single argument as used and returns true; if none of the flags is found it returns false and nothing is marked. It never returns an error: "not present" is a valid, expected outcome, not a failure. |
| `GetOptionsSize` | `func(flags []string) int` | GetOptionsSize counts how many arguments in Args equal one of the given flag spellings (e.g. []string{"-o", "--output"}), regardless of whether they have already been marked used. It never mutates Used — call it before looping over occurrence indices 0..size-1 with GetStringOption (or a typed variant) to read every occurrence of a repeatable option. |
| `GetKeyValuesSize` | `func(prefixes []string) int` | GetKeyValuesSize counts how many arguments in Args start with one of the given key=value prefixes (e.g. []string{"user=", "username="} — the separator is part of the prefix), regardless of Used. Like GetOptionsSize, it never mutates Used; pair it with GetStringKeyValues (or a typed variant) to iterate every match. |
| `GetStringOption` | `func(flags []string, occurrence int) (string, error)` | GetStringOption finds the occurrence-th (0-based) argument that equals one of the given flag spellings, then returns the argument immediately following it as the option's value. It marks both the flag and its value as used. It errors when occurrence is out of range for the number of matches (see GetOptionsSize), or when the matched flag is the last argument and has no following value to return. |
| `GetIntOption` | `func(flags []string, occurrence int) (int, error)` | GetIntOption behaves exactly like GetStringOption, additionally parsing the option's value as a base-10 integer. It returns every failure GetStringOption returns, plus a parse error when the value is not a valid integer. |
| `GetDoubleOption` | `func(flags []string, occurrence int) (float64, error)` | GetDoubleOption behaves exactly like GetStringOption, additionally parsing the option's value as a 64-bit floating-point number, erroring when the value is not a valid number. |
| `GetTimestampOption` | `func(flags []string, occurrence int) (int64, error)` | GetTimestampOption behaves exactly like GetStringOption, additionally parsing the option's value as an RFC 3339 timestamp and reporting it as nanoseconds since the Unix epoch, UTC. |
| `GetStringArg` | `func(index int) (string, error)` | GetStringArg returns the argument at the given absolute index of Args — the same index numbering as the raw command line (0 is the first argument after the program name), independent of which arguments have already been read. It marks that index as used, and errors when index is negative or beyond the end of Args. |
| `GetIntArg` | `func(index int) (int, error)` | GetIntArg behaves exactly like GetStringArg, additionally parsing the argument as a base-10 integer. |
| `GetDoubleArg` | `func(index int) (float64, error)` | GetDoubleArg behaves exactly like GetStringArg, additionally parsing the argument as a 64-bit floating-point number. |
| `GetTimestampArg` | `func(index int) (int64, error)` | GetTimestampArg behaves exactly like GetStringArg, additionally parsing the argument as an RFC 3339 timestamp and reporting it as nanoseconds since the Unix epoch, UTC. |
| `GetNextStringArg` | `func() (string, error)` | GetNextStringArg returns the first argument in Args, in order, whose Used entry is still false, and marks it used. This is the core of the Unused Mechanic: after every flag and option a program expects has been read with the functions above, whatever remains unread is exactly the leftover positional arguments — call this repeatedly to drain them in order. It errors when every argument has already been used. |
| `GetNextIntArg` | `func() (int, error)` | GetNextIntArg behaves exactly like GetNextStringArg, additionally parsing the argument as a base-10 integer. |
| `GetNextDoubleArg` | `func() (float64, error)` | GetNextDoubleArg behaves exactly like GetNextStringArg, additionally parsing the argument as a 64-bit floating-point number. |
| `GetNextTimestampArg` | `func() (int64, error)` | GetNextTimestampArg behaves exactly like GetNextStringArg, additionally parsing the argument as an RFC 3339 timestamp and reporting it as nanoseconds since the Unix epoch, UTC. |
| `GetStringKeyValues` | `func(prefixes []string, occurrence int) (string, error)` | GetStringKeyValues finds the occurrence-th (0-based) argument that starts with one of the given key=value prefixes (the separator is part of the prefix, e.g. "username="), then returns the text after the matched prefix as the value. It marks that argument as used. It errors when occurrence is out of range for the number of matches (see GetKeyValuesSize), or when the matched argument's value portion is empty (e.g. a bare "username=" with nothing after it). |
| `GetIntKeyValues` | `func(prefixes []string, occurrence int) (int, error)` | GetIntKeyValues behaves exactly like GetStringKeyValues, additionally parsing the value portion as a base-10 integer. |
| `GetDoubleKeyValues` | `func(prefixes []string, occurrence int) (float64, error)` | GetDoubleKeyValues behaves exactly like GetStringKeyValues, additionally parsing the value portion as a 64-bit floating-point number. |
| `GetTimestampKeyValues` | `func(prefixes []string, occurrence int) (int64, error)` | GetTimestampKeyValues behaves exactly like GetStringKeyValues, additionally parsing the value portion as an RFC 3339 timestamp and reporting it as nanoseconds since the Unix epoch, UTC. |

## `sandbox/api/info.go`

### `Info`

Info is the contract reporting the library's own identity, carried by the Sandbox as the field of the same name. Both values are compile-time constants of sandbox/internal/config, generated from AgnosConfig/project.yaml, so a release bump is a one-line edit touching no logic — and a caller can report which Verb it linked against without importing anything but this package.

| Field | Type | Description |
| --- | --- | --- |
| `Name` | `func() string` | Name is the library's name, "Verb". |
| `Version` | `func() string` | Version is the release the caller linked against, in the "v0.0.0" spelling the repository tags with. |

# Dependency contracts

`deps.Deps` has one field per directory of `sandbox/deps/`, named by title-casing it. Each
field is that package's `Sandbox` struct, filled by `adapters/libs/<name>.Bind(&deps)`.

## `deps.Std`

`sandbox/deps/std`

### `Sandbox`

Sandbox is the runtime library injected whole as the Deps.Std field.

| Field | Type | Description |
| --- | --- | --- |
| `Now` | `func() int64` | Now returns the current wall-clock time as nanoseconds since the Unix epoch, UTC. The sandbox may not name a `time.Time`, so an instant crosses this boundary as a plain integer. |
| `Printf` | `func(format string, a ...any) (n int, err error)` | Printf writes one formatted message to standard output. It carries the command's result — the data a script would read — so it is never silenced. |
| `Log` | `func(format string, a ...any) (n int, err error)` | Log writes one formatted progress message to standard error. It is the channel every "… started with path …" notice goes through, so a caller can keep stdout free of log noise, and it is what --quiet turns off. |
| `Error` | `func(format string, a ...any) (n int, err error)` | Error writes one formatted message to standard error. |
| `Errorf` | `func(format string, a ...any) error` | Errorf formats an error message and returns it as an error. |
| `Sprintf` | `func(format string, a ...any) string` | Sprintf formats a message and returns it as a string. It is the one formatting entry point the sandbox has: every string it builds out of values rather than out of concatenation goes through here. |
| `Goos` | `func() string` | Goos is the name of the operating system the process runs on, in the spelling the Go toolchain uses ("darwin", "linux", "windows", …). |

## `deps.Stringsdeps`

`sandbox/deps/stringsdeps`

### `Sandbox`

Sandbox is the text library injected whole as the Deps.Stringsdeps field. The first group of fields is string manipulation, the second is conversion between strings and numbers.

| Field | Type | Description |
| --- | --- | --- |
| `TrimSpace` | `func(s string) string` | TrimSpace returns s with leading and trailing white space removed. |
| `Trim` | `func(s string, cutset string) string` | Trim returns s with every leading and trailing character contained in cutset removed. |
| `TrimLeft` | `func(s string, cutset string) string` | TrimLeft returns s with every leading character contained in cutset removed. |
| `TrimRight` | `func(s string, cutset string) string` | TrimRight returns s with every trailing character contained in cutset removed. |
| `TrimPrefix` | `func(s string, prefix string) string` | TrimPrefix returns s without the given leading prefix. When s does not start with prefix, s is returned unchanged. |
| `TrimSuffix` | `func(s string, suffix string) string` | TrimSuffix returns s without the given trailing suffix. When s does not end with suffix, s is returned unchanged. |
| `HasPrefix` | `func(s string, prefix string) bool` | HasPrefix reports whether s begins with prefix. |
| `HasSuffix` | `func(s string, suffix string) bool` | HasSuffix reports whether s ends with suffix. |
| `Contains` | `func(s string, substr string) bool` | Contains reports whether substr is within s. |
| `ContainsAny` | `func(s string, chars string) bool` | ContainsAny reports whether any character of chars is within s. |
| `LastIndex` | `func(s string, substr string) int` | LastIndex returns the index of the last instance of substr in s, or -1 when substr is absent. |
| `Count` | `func(s string, substr string) int` | Count returns the number of non-overlapping instances of substr in s. When substr is empty it returns one plus the number of runes in s. |
| `Split` | `func(s string, sep string) []string` | Split slices s into every substring separated by sep. |
| `Join` | `func(elems []string, sep string) string` | Join concatenates elems, placing sep between consecutive elements. |
| `Fields` | `func(s string) []string` | Fields slices s around each run of white space, returning the substrings between them. |
| `FieldsFunc` | `func(s string, f func(rune) bool) []string` | FieldsFunc slices s at each run of runes satisfying f, returning the substrings between them. |
| `Repeat` | `func(s string, count int) string` | Repeat returns count copies of s concatenated. |
| `ReplaceAll` | `func(s string, old string, new string) string` | ReplaceAll returns s with every non-overlapping instance of old replaced by new. |
| `ToUpper` | `func(s string) string` | ToUpper returns s with every letter mapped to its upper case. |
| `ToLower` | `func(s string) string` | ToLower returns s with every letter mapped to its lower case. |
| `Quote` | `func(s string) string` | Quote returns s as a double-quoted Go string literal, escaping what the Go syntax requires. |
| `MatchPattern` | `func(pattern string, s string) (bool, error)` | MatchPattern reports whether s is matched by the regular expression pattern, and errors when the pattern itself does not compile. It is the one matching primitive the sandbox has: `regexp` lives on the adapter side like every other standard package. |
| `Atoi` | `func(s string) (int, error)` | Atoi parses s as a decimal integer. The error reports a string that is not one. |
| `ParseInt` | `func(s string, base int, bit_size int) (int64, error)` | ParseInt parses s as an integer in the given base with the given bit size. The error reports a string that is not one. |
| `ParseFloat` | `func(s string, bit_size int) (float64, error)` | ParseFloat parses s as a floating-point number of the given bit size. The error reports a string that is not one. |
| `FormatInt` | `func(value int64, base int) string` | FormatInt returns the string representation of value in the given base. |
| `FormatFloat` | `func(value float64, format byte, precision int, bit_size int) string` | FormatFloat returns the string representation of value, formatted according to the format byte, the precision and the bit size — the same three controls the standard library takes. |
