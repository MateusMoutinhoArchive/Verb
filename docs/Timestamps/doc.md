# Timestamps

`GetTimestampOption`, `GetTimestampArg`, `GetNextTimestampArg` and `GetTimestampKeyValues`
read RFC 3339 text and report **nanoseconds since the Unix epoch, UTC**, as an `int64`.

## Why an int64 and not a time.Time

Two rules meet here. `sandbox/` may import only `sandbox/` packages, so `time` cannot be
named inside it ([Rules](../Rules/doc.md#layers)); and every type of `sandbox/api/` has to
be convertible, so a consumer can copy the contract into a deps package of its own — a type
from another package is not. `sandbox/internal/rfc3339` therefore implements the grammar in
arithmetic alone, with no import but the api.

A caller that wants a `time.Time` writes one line on its own side:

```go
at := time.Unix(0, nanos).UTC()
```

## The grammar

```
2006-01-02T15:04:05Z07:00
```

| Accepted | Note |
|---|---|
| `2024-01-02T15:04:05Z` | UTC |
| `2024-01-02T12:04:05-03:00` | the offset is applied, so this is the same instant |
| `2024-01-02T15:04:05.25Z` | a fractional second, 1 to 9 digits |
| `2024-01-02t15:04:05z` | lowercase `t` and `z`, which the RFC allows |
| `2024-02-29T00:00:00Z` | a real leap day |

| Refused | Because |
|---|---|
| `2023-02-29T00:00:00Z` | 2023 is not a leap year — the calendar is real, not a format check |
| `2024-04-31T00:00:00Z` | April has 30 days |
| `2024-01-02T24:00:00Z` | hour, minute and second are `00-23`, `00-59`, `00-59` |
| `2024-01-02T15:04:05+0300` | the zone is `Z` or `±HH:MM`, colon included |
| `2024-01-02 15:04:05Z` | the date and the time are joined by `T` |
| `2024-01-02T15:04:05` | the zone is not optional |
| `1500-01-01T00:00:00Z` | outside the range below |

## Two divergences from `time.Parse`

Everything else matches the standard library exactly, instant for instant.

- **Range.** An `int64` of nanoseconds covers roughly **1678-01-01 to 2262-04-11**. An
  instant outside it is an error here and a `time.Time` there.
- **Zone offset.** RFC 3339 caps it at `23:59`; `time.Parse` accepts `+24:00` and `-00:60`
  and this does not.

A fractional second past the ninth digit is read and dropped, the way `time.Parse` reads it.

Every row above is pinned by [`examples/lib/timestamps`](../../examples/lib/timestamps/example.go),
whose golden is the regression check: the sandbox is closed to `testing`, so behaviour is
asserted by example and `exec-test`, never by a `_test.go` under `sandbox/`.
