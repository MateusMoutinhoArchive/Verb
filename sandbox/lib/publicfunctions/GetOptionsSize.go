package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetOptionsSizeFactory fills api.Lib.GetOptionsSize with a closure counting
// the arguments matching one of flags. It never mutates Used — it only
// counts, so the caller can then loop over occurrence indices 0..size-1.
func GetOptionsSizeFactory(l *api.Lib) func(flags []string) int {
	return func(flags []string) int {
		return argv.CountFlags(l, flags)
	}
}
