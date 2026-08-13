package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetIntKeyValuesFactory fills api.Lib.GetIntKeyValues with a closure
// behaving exactly like api.Lib.GetStringKeyValues, additionally parsing the
// value it finds as a base-10 integer.
func GetIntKeyValuesFactory(l *api.Lib) func(prefixes []string, occurrence int) (int, error) {
	return func(prefixes []string, occurrence int) (int, error) {
		s, err := argv.KeyValuesValue(l, prefixes, occurrence)
		if err != nil {
			return 0, err
		}
		return argv.ParseInt(s)
	}
}
