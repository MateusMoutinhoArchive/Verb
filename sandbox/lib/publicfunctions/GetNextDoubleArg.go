package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetNextDoubleArgFactory fills api.Lib.GetNextDoubleArg with a closure
// behaving exactly like api.Lib.GetNextStringArg, additionally parsing the
// value it finds as a 64-bit floating-point number.
func GetNextDoubleArgFactory(l *api.Lib) func() (float64, error) {
	return func() (float64, error) {
		s, err := argv.NextArgValue(l)
		if err != nil {
			return 0, err
		}
		return argv.ParseDouble(s)
	}
}
