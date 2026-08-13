package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetStringOptionFactory fills api.Lib.GetStringOption with a closure
// returning the value following the occurrence-th argument matching one of
// flags, marking both the flag and its value as used.
func GetStringOptionFactory(l *api.Lib) func(flags []string, occurrence int) (string, error) {
	return func(flags []string, occurrence int) (string, error) {
		return argv.OptionValue(l, flags, occurrence)
	}
}
