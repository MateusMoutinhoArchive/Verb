package sandbox

import (
	api "github.com/MateusMoutinhoOrg/Verb/sandbox/api"
	deps "github.com/MateusMoutinhoOrg/Verb/sandbox/deps"
	argv "github.com/MateusMoutinhoOrg/Verb/sandbox/internal/argv"
	info "github.com/MateusMoutinhoOrg/Verb/sandbox/internal/info"
)

func New(deps *deps.Deps) *api.Sandbox {
	self := api.Sandbox{Deps: deps}

	self.Argv = argv.NewArgv(&self)
	self.Info = info.NewInfo(&self)

	return &self
}
