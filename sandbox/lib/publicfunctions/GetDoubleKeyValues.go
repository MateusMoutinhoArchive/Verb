package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetDoubleKeyValuesFactory fills api.Lib.GetDoubleKeyValues with a closure
// behaving exactly like api.Lib.GetStringKeyValues, additionally parsing the
// value it finds as a 64-bit floating-point number.
func GetDoubleKeyValuesFactory(l *api.Lib) func(prefixes []string, occurrence int) (float64, error) {
	return func(prefixes []string, occurrence int) (float64, error) {
		s, err := argv.KeyValuesValue(l, prefixes, occurrence)
		if err != nil {
			return 0, err
		}
		return argv.ParseDouble(s)
	}
}
