package lib

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib/publicfunctions"
)

// New builds the api.Lib entry point: it stores args on Args, allocates a
// same-length Used tracking slice, and runs every lib factory over it to fill
// its function fields. Adding a function field to api.Lib means adding its
// factory call here — an unlisted field stays nil and panics on first call.
func New(args []string) api.Lib {
	l := api.Lib{
		Args: args,
		Used: make([]bool, len(args)),
	}

	l.IsPresent = publicfunctions.IsPresentFactory(&l)
	l.GetOptionsSize = publicfunctions.GetOptionsSizeFactory(&l)
	l.GetKeyValuesSize = publicfunctions.GetKeyValuesSizeFactory(&l)

	l.GetStringOption = publicfunctions.GetStringOptionFactory(&l)
	l.GetIntOption = publicfunctions.GetIntOptionFactory(&l)
	l.GetDoubleOption = publicfunctions.GetDoubleOptionFactory(&l)
	l.GetTimestampOption = publicfunctions.GetTimestampOptionFactory(&l)

	l.GetStringArg = publicfunctions.GetStringArgFactory(&l)
	l.GetIntArg = publicfunctions.GetIntArgFactory(&l)
	l.GetDoubleArg = publicfunctions.GetDoubleArgFactory(&l)
	l.GetTimestampArg = publicfunctions.GetTimestampArgFactory(&l)

	l.GetNextStringArg = publicfunctions.GetNextStringArgFactory(&l)
	l.GetNextIntArg = publicfunctions.GetNextIntArgFactory(&l)
	l.GetNextDoubleArg = publicfunctions.GetNextDoubleArgFactory(&l)
	l.GetNextTimestampArg = publicfunctions.GetNextTimestampArgFactory(&l)

	l.GetStringKeyValues = publicfunctions.GetStringKeyValuesFactory(&l)
	l.GetIntKeyValues = publicfunctions.GetIntKeyValuesFactory(&l)
	l.GetDoubleKeyValues = publicfunctions.GetDoubleKeyValuesFactory(&l)
	l.GetTimestampKeyValues = publicfunctions.GetTimestampKeyValuesFactory(&l)

	return l
}
