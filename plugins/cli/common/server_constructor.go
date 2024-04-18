package common

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/ikhfa/goxgen/plugins/cli/settings"
)

type Constructor func(settings *settings.EnvironmentSettings) (*handler.Server, error)
