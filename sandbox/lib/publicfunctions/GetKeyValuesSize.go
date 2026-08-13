package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetKeyValuesSizeFactory fills api.Lib.GetKeyValuesSize with a closure
// counting the arguments starting with one of prefixes. Like GetOptionsSize
// it never mutates Used — it only counts.
func GetKeyValuesSizeFactory(l *api.Lib) func(prefixes []string) int {
	return func(prefixes []string) int {
		return argv.CountPrefixes(l, prefixes)
	}
}
