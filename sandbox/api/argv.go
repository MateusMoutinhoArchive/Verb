package api

// This file is the whole of Verb's parsing surface: the Argv contract that
// reaches the Sandbox as the field of the same name, and the Parser it hands
// back.
//
// Every type here is a struct — never an interface. Parser carries behaviour,
// so it holds it as function fields, filled by sandbox/internal/argv at the
// moment the parser is built. That is what makes Verb cheap to embed: a
// consumer that restates this shape in a contract of its own fills it by
// assigning the fields straight across, with no wrapper type in between.

// Argv is the contract that turns an argument vector into a Parser, carried
// by the Sandbox as the field of the same name. A parser is bound to one
// argument vector and carries the read state of it, so it is built per call
// rather than once at wiring time — which is why this contract holds a
// constructor and nothing else.
type Argv struct {
	// New builds a Parser bound to the given arguments, e.g. os.Args[1:].
	// The slice is kept as it stands and is never written to; every
	// argument starts out unread. Passing nil or an empty slice is valid
	// and yields a parser that matches nothing.
	New func(args []string) Parser
}

// Parser is an argument-vector (argv) parser: the argv it was built over, the
// read state of it, and one function field per way of reading a value out of
// it. Every argument starts unread and every reader marks what it matched, so
// whatever is left over is exactly the positional arguments nothing asked for
// — docs/UnusedMechanic holds who marks what, and in what order to read.
// Each reader family (Option, Arg, NextArg, KeyValues) comes in four types:
// String (the raw text), Int (base-10), Double (a float64) and Timestamp (RFC
// 3339 text reported as nanoseconds since the Unix epoch, UTC —
// docs/Timestamps). A Get* function returns an error rather than a bool or a
// panic when it cannot produce a value; the two *Size functions and
// IsPresent return plain values, because counting and checking presence
// cannot fail.
type Parser struct {
	// Args is the argument vector being parsed, the same slice Argv.New was
	// called with. Every index-based function refers to positions in this
	// slice. Treat it as read-only: mutating it after construction leaves
	// Used out of sync and produces undefined matching behaviour.
	Args []string

	// Used tracks, index for index against Args, which arguments have
	// already been matched by a previous call. Used[i] is true once Args[i]
	// has been consumed by any Get* function or by IsPresent. It starts
	// entirely false and only grows more true over the parser's lifetime —
	// the Unused Mechanic reads it to find the next never-consumed
	// positional argument. Treat it as read-only.
	Used []bool

	// IsPresent reports whether any of the given flag spellings (e.g.
	// []string{"-q", "--quiet"}) occurs anywhere in the unread portion of
	// Args. On a match it marks that single argument as used and returns
	// true; if none of the flags is found it returns false and nothing is
	// marked. It never returns an error: "not present" is a valid, expected
	// outcome, not a failure.
	IsPresent func(flags []string) bool

	// GetOptionsSize counts how many arguments in Args equal one of the
	// given flag spellings (e.g. []string{"-o", "--output"}), regardless of
	// whether they have already been marked used. It never mutates Used —
	// call it before looping over occurrence indices 0..size-1 with
	// GetStringOption (or a typed variant) to read every occurrence of a
	// repeatable option.
	GetOptionsSize func(flags []string) int

	// GetKeyValuesSize counts how many arguments in Args start with one of
	// the given key=value prefixes (e.g. []string{"user=", "username="} —
	// the separator is part of the prefix), regardless of Used. Like
	// GetOptionsSize, it never mutates Used; pair it with
	// GetStringKeyValues (or a typed variant) to iterate every match.
	GetKeyValuesSize func(prefixes []string) int

	// GetStringOption finds the occurrence-th (0-based) argument that equals
	// one of the given flag spellings, then returns the argument immediately
	// following it as the option's value. It marks both the flag and its
	// value as used. It errors when occurrence is out of range for the
	// number of matches (see GetOptionsSize), or when the matched flag is
	// the last argument and has no following value to return.
	GetStringOption func(flags []string, occurrence int) (string, error)
	// GetIntOption behaves exactly like GetStringOption, additionally
	// parsing the option's value as a base-10 integer. It returns every
	// failure GetStringOption returns, plus a parse error when the value is
	// not a valid integer.
	GetIntOption func(flags []string, occurrence int) (int, error)
	// GetDoubleOption behaves exactly like GetStringOption, additionally
	// parsing the option's value as a 64-bit floating-point number,
	// erroring when the value is not a valid number.
	GetDoubleOption func(flags []string, occurrence int) (float64, error)
	// GetTimestampOption behaves exactly like GetStringOption, additionally
	// parsing the option's value as an RFC 3339 timestamp and reporting it
	// as nanoseconds since the Unix epoch, UTC.
	GetTimestampOption func(flags []string, occurrence int) (int64, error)

	// GetStringArg returns the argument at the given absolute index of Args
	// — the same index numbering as the raw command line (0 is the first
	// argument after the program name), independent of which arguments have
	// already been read. It marks that index as used, and errors when index
	// is negative or beyond the end of Args.
	GetStringArg func(index int) (string, error)
	// GetIntArg behaves exactly like GetStringArg, additionally parsing the
	// argument as a base-10 integer.
	GetIntArg func(index int) (int, error)
	// GetDoubleArg behaves exactly like GetStringArg, additionally parsing
	// the argument as a 64-bit floating-point number.
	GetDoubleArg func(index int) (float64, error)
	// GetTimestampArg behaves exactly like GetStringArg, additionally
	// parsing the argument as an RFC 3339 timestamp and reporting it as
	// nanoseconds since the Unix epoch, UTC.
	GetTimestampArg func(index int) (int64, error)

	// GetNextStringArg returns the first argument in Args, in order, whose
	// Used entry is still false, and marks it used. This is the core of the
	// Unused Mechanic: after every flag and option a program expects has
	// been read with the functions above, whatever remains unread is exactly
	// the leftover positional arguments — call this repeatedly to drain them
	// in order. It errors when every argument has already been used.
	GetNextStringArg func() (string, error)
	// GetNextIntArg behaves exactly like GetNextStringArg, additionally
	// parsing the argument as a base-10 integer.
	GetNextIntArg func() (int, error)
	// GetNextDoubleArg behaves exactly like GetNextStringArg, additionally
	// parsing the argument as a 64-bit floating-point number.
	GetNextDoubleArg func() (float64, error)
	// GetNextTimestampArg behaves exactly like GetNextStringArg,
	// additionally parsing the argument as an RFC 3339 timestamp and
	// reporting it as nanoseconds since the Unix epoch, UTC.
	GetNextTimestampArg func() (int64, error)

	// GetStringKeyValues finds the occurrence-th (0-based) argument that
	// starts with one of the given key=value prefixes (the separator is part
	// of the prefix, e.g. "username="), then returns the text after the
	// matched prefix as the value. It marks that argument as used. It errors
	// when occurrence is out of range for the number of matches (see
	// GetKeyValuesSize), or when the matched argument's value portion is
	// empty (e.g. a bare "username=" with nothing after it).
	GetStringKeyValues func(prefixes []string, occurrence int) (string, error)
	// GetIntKeyValues behaves exactly like GetStringKeyValues, additionally
	// parsing the value portion as a base-10 integer.
	GetIntKeyValues func(prefixes []string, occurrence int) (int, error)
	// GetDoubleKeyValues behaves exactly like GetStringKeyValues,
	// additionally parsing the value portion as a 64-bit floating-point
	// number.
	GetDoubleKeyValues func(prefixes []string, occurrence int) (float64, error)
	// GetTimestampKeyValues behaves exactly like GetStringKeyValues,
	// additionally parsing the value portion as an RFC 3339 timestamp and
	// reporting it as nanoseconds since the Unix epoch, UTC.
	GetTimestampKeyValues func(prefixes []string, occurrence int) (int64, error)
}
