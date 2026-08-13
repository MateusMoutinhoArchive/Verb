package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// IsPresentFactory fills api.Lib.IsPresent with a closure matching any
// still-unused argument against flags, marking it used on a hit. It never
// fails: "not present" is a valid outcome, reported as false.
func IsPresentFactory(l *api.Lib) func(flags []string) bool {
	return func(flags []string) bool {
		_, found := argv.FirstUnusedFlag(l, flags)
		return found
	}
}
