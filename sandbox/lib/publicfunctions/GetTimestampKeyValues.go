package publicfunctions

import (
	"time"

	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/argv"
)

// GetTimestampKeyValuesFactory fills api.Lib.GetTimestampKeyValues with a
// closure behaving exactly like api.Lib.GetStringKeyValues, additionally
// parsing the value it finds as an RFC 3339 timestamp.
func GetTimestampKeyValuesFactory(l *api.Lib) func(prefixes []string, occurrence int) (time.Time, error) {
	return func(prefixes []string, occurrence int) (time.Time, error) {
		s, err := argv.KeyValuesValue(l, prefixes, occurrence)
		if err != nil {
			return time.Time{}, err
		}
		return argv.ParseTimestamp(s)
	}
}
