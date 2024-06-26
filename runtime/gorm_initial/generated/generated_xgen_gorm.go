package generated

import (
	"fmt"
	"github.com/ikhfa/goxgen/graphql/sort"
	"github.com/ikhfa/goxgen/plugins/cli/settings"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewGormDB(sts *settings.EnvironmentSettings) (*gorm.DB, error) {
	dbDriver := sts.DatabaseDriver
	dbDsn := sts.DatabaseDsn
	var dialector gorm.Dialector
	switch dbDriver {
	case "sqlite":
		dialector = sqlite.Open(dbDsn)
	case "mysql":
		dialector = mysql.Open(dbDsn)
	case "postgres":
		dialector = postgres.Open(dbDsn)
	case "sqlserver":
		dialector = sqlserver.Open(dbDsn)
	case "clickhouse":
		dialector = clickhouse.Open(dbDsn)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", dbDriver)
	}

	// Open the database connection
	db, err := gorm.Open(
		dialector,
		&gorm.Config{
			FullSaveAssociations: true,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Migrate the schema
	err = db.AutoMigrate()
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}

func Paginate(paginationInput *XgenPaginationInput) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if paginationInput == nil {
			return db
		}

		page := paginationInput.Page
		if page <= 0 {
			page = 1
		}

		pageSize := paginationInput.Size
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func Sort(sort sort.Sortables) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, s := range sort.SortSqlStrings() {
			db = db.Order(fmt.Sprintf(s))
		}
		return db
	}
}
