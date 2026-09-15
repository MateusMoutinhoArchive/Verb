package standard

import (
	std "github.com/MateusMoutinhoOrg/Verb/adapters/libs/std"
	stringsdeps "github.com/MateusMoutinhoOrg/Verb/adapters/libs/stringsdeps"
	deps "github.com/MateusMoutinhoOrg/Verb/sandbox/deps"
)

func New() deps.Deps {
	deps := deps.Deps{}
	std.Bind(&deps)
	stringsdeps.Bind(&deps)
	return deps
}
