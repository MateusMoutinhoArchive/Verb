package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetIntOptionFactory fills api.Lib.GetIntOption with a closure behaving
// exactly like api.Lib.GetStringOption, additionally parsing the value it
// finds as a base-10 integer.
func GetIntOptionFactory(l *api.Lib) func(flags []string, occurrence int) (int, error) {
	return func(flags []string, occurrence int) (int, error) {
		s, err := argv.OptionValue(l, flags, occurrence)
		if err != nil {
			return 0, err
		}
		return argv.ParseInt(s)
	}
}
