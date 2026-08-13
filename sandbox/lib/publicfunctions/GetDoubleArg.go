package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetDoubleArgFactory fills api.Lib.GetDoubleArg with a closure behaving
// exactly like api.Lib.GetStringArg, additionally parsing the value it finds
// as a 64-bit floating-point number.
func GetDoubleArgFactory(l *api.Lib) func(index int) (float64, error) {
	return func(index int) (float64, error) {
		s, err := argv.ArgValue(l, index)
		if err != nil {
			return 0, err
		}
		return argv.ParseDouble(s)
	}
}
