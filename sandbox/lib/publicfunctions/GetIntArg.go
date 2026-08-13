package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetIntArgFactory fills api.Lib.GetIntArg with a closure behaving exactly
// like api.Lib.GetStringArg, additionally parsing the value it finds as a
// base-10 integer.
func GetIntArgFactory(l *api.Lib) func(index int) (int, error) {
	return func(index int) (int, error) {
		s, err := argv.ArgValue(l, index)
		if err != nil {
			return 0, err
		}
		return argv.ParseInt(s)
	}
}
