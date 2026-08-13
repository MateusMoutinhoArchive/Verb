package publicfunctions

import (
	"time"

	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetTimestampArgFactory fills api.Lib.GetTimestampArg with a closure
// behaving exactly like api.Lib.GetStringArg, additionally parsing the value
// it finds as an RFC 3339 timestamp.
func GetTimestampArgFactory(l *api.Lib) func(index int) (time.Time, error) {
	return func(index int) (time.Time, error) {
		s, err := argv.ArgValue(l, index)
		if err != nil {
			return time.Time{}, err
		}
		return argv.ParseTimestamp(s)
	}
}
