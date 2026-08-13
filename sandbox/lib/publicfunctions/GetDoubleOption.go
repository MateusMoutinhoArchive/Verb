package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetDoubleOptionFactory fills api.Lib.GetDoubleOption with a closure
// behaving exactly like api.Lib.GetStringOption, additionally parsing the
// value it finds as a 64-bit floating-point number.
func GetDoubleOptionFactory(l *api.Lib) func(flags []string, occurrence int) (float64, error) {
	return func(flags []string, occurrence int) (float64, error) {
		s, err := argv.OptionValue(l, flags, occurrence)
		if err != nil {
			return 0, err
		}
		return argv.ParseDouble(s)
	}
}
