package argv

import (
	api "github.com/MateusMoutinhoOrg/Verb/sandbox/api"
)

// Locating an argument and consuming it. Every reader below is written once
// and shared by the four typed variants of its family, so String, Int, Double
// and Timestamp differ in what they do to the text they were handed and in
// nothing else — a getter that matches marks the same indexes used whether
// the value then parses or not.

// MatchesFlag reports whether argument is spelled exactly like one of flags.
func MatchesFlag(sandbox *api.Sandbox, argument string, flags []string) bool {
	for _, flag := range flags {
		if argument == flag {
			return true
		}
	}
	return false
}

// MatchPrefix reports whether argument starts with one of the given key=value
// prefixes, returning the prefix that matched.
func MatchPrefix(sandbox *api.Sandbox, argument string, prefixes []string) (string, bool) {
	for _, prefix := range prefixes {
		if sandbox.Deps.Stringsdeps.HasPrefix(argument, prefix) {
			return prefix, true
		}
	}
	return "", false
}

// CountFlags counts how many arguments are spelled like one of flags, without
// touching parser.Used.
func CountFlags(sandbox *api.Sandbox, parser *api.Parser, flags []string) int {
	count := 0
	for _, argument := range parser.Args {
		if MatchesFlag(sandbox, argument, flags) {
			count++
		}
	}
	return count
}

// CountPrefixes counts how many arguments start with one of prefixes, without
// touching parser.Used.
func CountPrefixes(sandbox *api.Sandbox, parser *api.Parser, prefixes []string) int {
	count := 0
	for _, argument := range parser.Args {
		if _, matched := MatchPrefix(sandbox, argument, prefixes); matched {
			count++
		}
	}
	return count
}

// FirstUnusedFlag finds the first not-yet-used argument spelled like one of
// flags and marks it used, reporting whether there was one. It backs
// IsPresent: a flag carries no value, so the match itself is the answer.
func FirstUnusedFlag(sandbox *api.Sandbox, parser *api.Parser, flags []string) bool {
	for index, argument := range parser.Args {
		if parser.Used[index] {
			continue
		}
		if MatchesFlag(sandbox, argument, flags) {
			parser.Used[index] = true
			return true
		}
	}
	return false
}

// findOptionIndex returns the position of the occurrence-th argument spelled
// like one of flags, regardless of Used — occurrence numbering counts every
// match, so it does not shift as earlier calls consume arguments.
func findOptionIndex(sandbox *api.Sandbox, parser *api.Parser, flags []string, occurrence int) (int, error) {
	count := 0
	for index, argument := range parser.Args {
		if !MatchesFlag(sandbox, argument, flags) {
			continue
		}
		if count == occurrence {
			return index, nil
		}
		count++
	}
	return -1, sandbox.Deps.Std.Errorf("verb: option %v: occurrence %d not found (only %d present)", flags, occurrence, count)
}

// OptionValue backs the Option getters: it locates the occurrence-th flag,
// marks it and the argument after it used, and returns that argument.
func OptionValue(sandbox *api.Sandbox, parser *api.Parser, flags []string, occurrence int) (string, error) {
	index, err := findOptionIndex(sandbox, parser, flags, occurrence)
	if err != nil {
		return "", err
	}
	if index+1 >= len(parser.Args) {
		return "", sandbox.Deps.Std.Errorf("verb: option %v: flag %q has no following value", flags, parser.Args[index])
	}
	parser.Used[index] = true
	parser.Used[index+1] = true
	return parser.Args[index+1], nil
}

// ArgValue backs the Arg getters: it checks index against the argv, marks it
// used, and returns the argument standing at that absolute position.
func ArgValue(sandbox *api.Sandbox, parser *api.Parser, index int) (string, error) {
	if index < 0 || index >= len(parser.Args) {
		return "", sandbox.Deps.Std.Errorf("verb: arg index %d out of range (have %d arguments)", index, len(parser.Args))
	}
	parser.Used[index] = true
	return parser.Args[index], nil
}

// NextArgValue backs the NextArg getters — the Unused Mechanic: it finds the
// first argument nothing has read yet, marks it used, and returns it.
func NextArgValue(sandbox *api.Sandbox, parser *api.Parser) (string, error) {
	for index, used := range parser.Used {
		if used {
			continue
		}
		parser.Used[index] = true
		return parser.Args[index], nil
	}
	return "", sandbox.Deps.Std.Errorf("verb: no unused arguments remaining")
}

// KeyValuesValue backs the KeyValues getters: it locates the occurrence-th
// argument starting with one of prefixes, marks it used, and returns the text
// standing after the prefix that matched.
func KeyValuesValue(sandbox *api.Sandbox, parser *api.Parser, prefixes []string, occurrence int) (string, error) {
	count := 0
	for index, argument := range parser.Args {
		prefix, matched := MatchPrefix(sandbox, argument, prefixes)
		if !matched {
			continue
		}
		if count == occurrence {
			parser.Used[index] = true
			value := argument[len(prefix):]
			if value == "" {
				return "", sandbox.Deps.Std.Errorf("verb: key/value %v: occurrence %d has an empty value", prefixes, occurrence)
			}
			return value, nil
		}
		count++
	}
	return "", sandbox.Deps.Std.Errorf("verb: key/value %v: occurrence %d not found (only %d present)", prefixes, occurrence, count)
}
