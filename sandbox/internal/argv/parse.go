package argv

import (
	api "github.com/MateusMoutinhoOrg/Verb/sandbox/api"
	rfc3339 "github.com/MateusMoutinhoOrg/Verb/sandbox/internal/rfc3339"
)

// Turning the text a reader returned into a typed value. These are the whole
// of the difference between the four variants of a getter family, and each is
// called after the match has already been marked used — a malformed value is
// a value that was read, not one that was missed.

// ParseInt parses text as a base-10 integer, for the Int getters.
func ParseInt(sandbox *api.Sandbox, text string) (int, error) {
	value, err := sandbox.Deps.Stringsdeps.Atoi(text)
	if err != nil {
		return 0, sandbox.Deps.Std.Errorf("verb: %q is not a valid integer: %w", text, err)
	}
	return value, nil
}

// ParseDouble parses text as a 64-bit floating-point number, for the Double
// getters.
func ParseDouble(sandbox *api.Sandbox, text string) (float64, error) {
	value, err := sandbox.Deps.Stringsdeps.ParseFloat(text, 64)
	if err != nil {
		return 0, sandbox.Deps.Std.Errorf("verb: %q is not a valid number: %w", text, err)
	}
	return value, nil
}

// ParseTimestamp parses text as an RFC 3339 timestamp and reports it as
// nanoseconds since the Unix epoch, UTC, for the Timestamp getters.
func ParseTimestamp(sandbox *api.Sandbox, text string) (int64, error) {
	return rfc3339.ParseNano(sandbox, text)
}
