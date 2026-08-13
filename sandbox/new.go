package lib

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/sandbox/lib"
)

// New builds the api.Lib entry point from the argument vector to parse,
// e.g. os.Args[1:]. It delegates to the sandbox/lib constructor, which
// stores args on the struct and runs the factories over it, each of which
// fills one function field with a closure reading it.
func New(args []string) api.Lib {
	return lib.New(args)
}
