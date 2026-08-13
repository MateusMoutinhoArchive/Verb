package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetStringKeyValuesFactory fills api.Lib.GetStringKeyValues with a closure
// returning the text after the matched prefix of the occurrence-th argument
// starting with one of prefixes, marking it used.
func GetStringKeyValuesFactory(l *api.Lib) func(prefixes []string, occurrence int) (string, error) {
	return func(prefixes []string, occurrence int) (string, error) {
		return argv.KeyValuesValue(l, prefixes, occurrence)
	}
}
