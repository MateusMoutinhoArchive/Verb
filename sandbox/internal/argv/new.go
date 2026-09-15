package argv

import (
	api "github.com/MateusMoutinhoOrg/Verb/sandbox/api"
)

// One factory per function field of api.Parser, each returning a closure that
// reads the parser it was handed at call time rather than at build time —
// which is what lets the read state of the argv live in one place while
// twenty independent functions consume it.
//
// The four variants of a family share their reader and differ only in the
// parser they run over the text it returned, so every one of them marks the
// same indexes used whether the value then parses or not.

// IsPresentFactory fills api.Parser.IsPresent.
func IsPresentFactory(sandbox *api.Sandbox, parser *api.Parser) func(flags []string) bool {
	return func(flags []string) bool {
		return FirstUnusedFlag(sandbox, parser, flags)
	}
}

// GetOptionsSizeFactory fills api.Parser.GetOptionsSize.
func GetOptionsSizeFactory(sandbox *api.Sandbox, parser *api.Parser) func(flags []string) int {
	return func(flags []string) int {
		return CountFlags(sandbox, parser, flags)
	}
}

// GetKeyValuesSizeFactory fills api.Parser.GetKeyValuesSize.
func GetKeyValuesSizeFactory(sandbox *api.Sandbox, parser *api.Parser) func(prefixes []string) int {
	return func(prefixes []string) int {
		return CountPrefixes(sandbox, parser, prefixes)
	}
}

// GetStringOptionFactory fills api.Parser.GetStringOption.
func GetStringOptionFactory(sandbox *api.Sandbox, parser *api.Parser) func(flags []string, occurrence int) (string, error) {
	return func(flags []string, occurrence int) (string, error) {
		return OptionValue(sandbox, parser, flags, occurrence)
	}
}

// GetIntOptionFactory fills api.Parser.GetIntOption.
func GetIntOptionFactory(sandbox *api.Sandbox, parser *api.Parser) func(flags []string, occurrence int) (int, error) {
	return func(flags []string, occurrence int) (int, error) {
		text, err := OptionValue(sandbox, parser, flags, occurrence)
		if err != nil {
			return 0, err
		}
		return ParseInt(sandbox, text)
	}
}

// GetDoubleOptionFactory fills api.Parser.GetDoubleOption.
func GetDoubleOptionFactory(sandbox *api.Sandbox, parser *api.Parser) func(flags []string, occurrence int) (float64, error) {
	return func(flags []string, occurrence int) (float64, error) {
		text, err := OptionValue(sandbox, parser, flags, occurrence)
		if err != nil {
			return 0, err
		}
		return ParseDouble(sandbox, text)
	}
}

// GetTimestampOptionFactory fills api.Parser.GetTimestampOption.
func GetTimestampOptionFactory(sandbox *api.Sandbox, parser *api.Parser) func(flags []string, occurrence int) (int64, error) {
	return func(flags []string, occurrence int) (int64, error) {
		text, err := OptionValue(sandbox, parser, flags, occurrence)
		if err != nil {
			return 0, err
		}
		return ParseTimestamp(sandbox, text)
	}
}

// GetStringArgFactory fills api.Parser.GetStringArg.
func GetStringArgFactory(sandbox *api.Sandbox, parser *api.Parser) func(index int) (string, error) {
	return func(index int) (string, error) {
		return ArgValue(sandbox, parser, index)
	}
}

// GetIntArgFactory fills api.Parser.GetIntArg.
func GetIntArgFactory(sandbox *api.Sandbox, parser *api.Parser) func(index int) (int, error) {
	return func(index int) (int, error) {
		text, err := ArgValue(sandbox, parser, index)
		if err != nil {
			return 0, err
		}
		return ParseInt(sandbox, text)
	}
}

// GetDoubleArgFactory fills api.Parser.GetDoubleArg.
func GetDoubleArgFactory(sandbox *api.Sandbox, parser *api.Parser) func(index int) (float64, error) {
	return func(index int) (float64, error) {
		text, err := ArgValue(sandbox, parser, index)
		if err != nil {
			return 0, err
		}
		return ParseDouble(sandbox, text)
	}
}

// GetTimestampArgFactory fills api.Parser.GetTimestampArg.
func GetTimestampArgFactory(sandbox *api.Sandbox, parser *api.Parser) func(index int) (int64, error) {
	return func(index int) (int64, error) {
		text, err := ArgValue(sandbox, parser, index)
		if err != nil {
			return 0, err
		}
		return ParseTimestamp(sandbox, text)
	}
}

// GetNextStringArgFactory fills api.Parser.GetNextStringArg.
func GetNextStringArgFactory(sandbox *api.Sandbox, parser *api.Parser) func() (string, error) {
	return func() (string, error) {
		return NextArgValue(sandbox, parser)
	}
}

// GetNextIntArgFactory fills api.Parser.GetNextIntArg.
func GetNextIntArgFactory(sandbox *api.Sandbox, parser *api.Parser) func() (int, error) {
	return func() (int, error) {
		text, err := NextArgValue(sandbox, parser)
		if err != nil {
			return 0, err
		}
		return ParseInt(sandbox, text)
	}
}

// GetNextDoubleArgFactory fills api.Parser.GetNextDoubleArg.
func GetNextDoubleArgFactory(sandbox *api.Sandbox, parser *api.Parser) func() (float64, error) {
	return func() (float64, error) {
		text, err := NextArgValue(sandbox, parser)
		if err != nil {
			return 0, err
		}
		return ParseDouble(sandbox, text)
	}
}

// GetNextTimestampArgFactory fills api.Parser.GetNextTimestampArg.
func GetNextTimestampArgFactory(sandbox *api.Sandbox, parser *api.Parser) func() (int64, error) {
	return func() (int64, error) {
		text, err := NextArgValue(sandbox, parser)
		if err != nil {
			return 0, err
		}
		return ParseTimestamp(sandbox, text)
	}
}

// GetStringKeyValuesFactory fills api.Parser.GetStringKeyValues.
func GetStringKeyValuesFactory(sandbox *api.Sandbox, parser *api.Parser) func(prefixes []string, occurrence int) (string, error) {
	return func(prefixes []string, occurrence int) (string, error) {
		return KeyValuesValue(sandbox, parser, prefixes, occurrence)
	}
}

// GetIntKeyValuesFactory fills api.Parser.GetIntKeyValues.
func GetIntKeyValuesFactory(sandbox *api.Sandbox, parser *api.Parser) func(prefixes []string, occurrence int) (int, error) {
	return func(prefixes []string, occurrence int) (int, error) {
		text, err := KeyValuesValue(sandbox, parser, prefixes, occurrence)
		if err != nil {
			return 0, err
		}
		return ParseInt(sandbox, text)
	}
}

// GetDoubleKeyValuesFactory fills api.Parser.GetDoubleKeyValues.
func GetDoubleKeyValuesFactory(sandbox *api.Sandbox, parser *api.Parser) func(prefixes []string, occurrence int) (float64, error) {
	return func(prefixes []string, occurrence int) (float64, error) {
		text, err := KeyValuesValue(sandbox, parser, prefixes, occurrence)
		if err != nil {
			return 0, err
		}
		return ParseDouble(sandbox, text)
	}
}

// GetTimestampKeyValuesFactory fills api.Parser.GetTimestampKeyValues.
func GetTimestampKeyValuesFactory(sandbox *api.Sandbox, parser *api.Parser) func(prefixes []string, occurrence int) (int64, error) {
	return func(prefixes []string, occurrence int) (int64, error) {
		text, err := KeyValuesValue(sandbox, parser, prefixes, occurrence)
		if err != nil {
			return 0, err
		}
		return ParseTimestamp(sandbox, text)
	}
}

// NewParser builds one api.Parser bound to args: it keeps the argv as it
// stands, allocates the same-length Used slice tracking what has been read
// out of it, and runs every factory over it to fill its function fields.
// Adding a function field to api.Parser means adding its factory call here —
// an unlisted field stays nil and panics on first call.
//
// The value returned is a copy of the struct the closures hold a pointer to,
// which is sound because what they write to is Used: a slice header the copy
// shares the backing array of. Neither slice is ever regrown, so a caller
// reading parser.Used sees exactly what the getters have marked.
func NewParser(sandbox *api.Sandbox, args []string) api.Parser {
	parser := api.Parser{
		Args: args,
		Used: make([]bool, len(args)),
	}

	parser.IsPresent = IsPresentFactory(sandbox, &parser)
	parser.GetOptionsSize = GetOptionsSizeFactory(sandbox, &parser)
	parser.GetKeyValuesSize = GetKeyValuesSizeFactory(sandbox, &parser)

	parser.GetStringOption = GetStringOptionFactory(sandbox, &parser)
	parser.GetIntOption = GetIntOptionFactory(sandbox, &parser)
	parser.GetDoubleOption = GetDoubleOptionFactory(sandbox, &parser)
	parser.GetTimestampOption = GetTimestampOptionFactory(sandbox, &parser)

	parser.GetStringArg = GetStringArgFactory(sandbox, &parser)
	parser.GetIntArg = GetIntArgFactory(sandbox, &parser)
	parser.GetDoubleArg = GetDoubleArgFactory(sandbox, &parser)
	parser.GetTimestampArg = GetTimestampArgFactory(sandbox, &parser)

	parser.GetNextStringArg = GetNextStringArgFactory(sandbox, &parser)
	parser.GetNextIntArg = GetNextIntArgFactory(sandbox, &parser)
	parser.GetNextDoubleArg = GetNextDoubleArgFactory(sandbox, &parser)
	parser.GetNextTimestampArg = GetNextTimestampArgFactory(sandbox, &parser)

	parser.GetStringKeyValues = GetStringKeyValuesFactory(sandbox, &parser)
	parser.GetIntKeyValues = GetIntKeyValuesFactory(sandbox, &parser)
	parser.GetDoubleKeyValues = GetDoubleKeyValuesFactory(sandbox, &parser)
	parser.GetTimestampKeyValues = GetTimestampKeyValuesFactory(sandbox, &parser)

	return parser
}

// NewFactory fills api.Argv.New.
func NewFactory(sandbox *api.Sandbox, argv *api.Argv) func(args []string) api.Parser {
	return func(args []string) api.Parser {
		return NewParser(sandbox, args)
	}
}

// NewArgv builds the api.Argv contract, running every factory over it to fill
// its function fields. Adding a function field to api.Argv means adding its
// factory call here.
func NewArgv(sandbox *api.Sandbox) api.Argv {
	argv := api.Argv{}
	argv.New = NewFactory(sandbox, &argv)
	return argv
}
