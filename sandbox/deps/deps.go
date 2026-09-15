package deps

import (
	std "github.com/MateusMoutinhoOrg/Verb/sandbox/deps/std"
	stringsdeps "github.com/MateusMoutinhoOrg/Verb/sandbox/deps/stringsdeps"
)

// Deps is every capability the sandbox needs from the outside world, one field
// per sub-contract directory of sandbox/deps/. An adapter fills the fields; the
// sandbox only calls them, which is what keeps it free of OS packages.
type Deps struct {
	Std         std.Sandbox
	Stringsdeps stringsdeps.Sandbox
}
