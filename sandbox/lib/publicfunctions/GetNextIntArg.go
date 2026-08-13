package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetNextIntArgFactory fills api.Lib.GetNextIntArg with a closure behaving
// exactly like api.Lib.GetNextStringArg, additionally parsing the value it
// finds as a base-10 integer.
func GetNextIntArgFactory(l *api.Lib) func() (int, error) {
	return func() (int, error) {
		s, err := argv.NextArgValue(l)
		if err != nil {
			return 0, err
		}
		return argv.ParseInt(s)
	}
}
