package publicfunctions

import (
	"time"

	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetTimestampOptionFactory fills api.Lib.GetTimestampOption with a closure
// behaving exactly like api.Lib.GetStringOption, additionally parsing the
// value it finds as an RFC 3339 timestamp.
func GetTimestampOptionFactory(l *api.Lib) func(flags []string, occurrence int) (time.Time, error) {
	return func(flags []string, occurrence int) (time.Time, error) {
		s, err := argv.OptionValue(l, flags, occurrence)
		if err != nil {
			return time.Time{}, err
		}
		return argv.ParseTimestamp(s)
	}
}
