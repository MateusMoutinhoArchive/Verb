package config

import "time"

// TimestampLayout is the layout every Timestamp getter of api.Lib parses its
// value with — GetTimestampOption, GetTimestampArg, GetNextTimestampArg and
// GetTimestampKeyValues alike. It lives in a file of its own so changing the
// accepted timestamp format is a one-line edit touching no parsing logic.
const TimestampLayout = time.RFC3339
