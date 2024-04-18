package simple_initial

import (
	"github.com/ikhfa/goxgen/plugins/cli/settings"
)

type Resolver struct{}

func NewResolver(sts *settings.EnvironmentSettings) (*Resolver, error) {
	return &Resolver{}, nil
}
