package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetStringArgFactory fills api.Lib.GetStringArg with a closure returning
// the argument at the given absolute index of Args, marking it used.
func GetStringArgFactory(l *api.Lib) func(index int) (string, error) {
	return func(index int) (string, error) {
		return argv.ArgValue(l, index)
	}
}
