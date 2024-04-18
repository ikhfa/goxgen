package gorm_example

import (
	"embed"
	"fmt"
	"github.com/ikhfa/goxgen/cmd/internal/integration/gorm_example/generated"
	"github.com/ikhfa/goxgen/plugins/cli/settings"
	"gorm.io/gorm"
)

//go:embed tests/*
var TestsFS embed.FS

type Resolver struct {
	DB *gorm.DB
}

func NewResolver(sts *settings.EnvironmentSettings) (*Resolver, error) {
	r := &Resolver{}
	db, err := generated.NewGormDB(sts)
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm db: %w", err)
	}
	r.DB = db

	return r, nil
}
