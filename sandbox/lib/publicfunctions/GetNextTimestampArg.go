package publicfunctions

import (
	"time"

	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetNextTimestampArgFactory fills api.Lib.GetNextTimestampArg with a
// closure behaving exactly like api.Lib.GetNextStringArg, additionally
// parsing the value it finds as an RFC 3339 timestamp.
func GetNextTimestampArgFactory(l *api.Lib) func() (time.Time, error) {
	return func() (time.Time, error) {
		s, err := argv.NextArgValue(l)
		if err != nil {
			return time.Time{}, err
		}
		return argv.ParseTimestamp(s)
	}
}
