package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetNextStringArgFactory fills api.Lib.GetNextStringArg with a closure
// returning the first still-unused argument in order, marking it used — the
// Unused Mechanic.
func GetNextStringArgFactory(l *api.Lib) func() (string, error) {
	return func() (string, error) {
		return argv.NextArgValue(l)
	}
}
