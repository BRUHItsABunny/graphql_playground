package graph

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
	"os"
)

func Connect() (*bun.DB, error) {
	// connect
	dsn := os.Getenv("POSTGRES_URL")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	// find the migrations
	var migrationsToExec = migrate.NewMigrations()
	err := migrationsToExec.Discover(os.DirFS("./_resources/migrations"))
	if err != nil {
		return nil, errors.Wrap(err, "Error loading migrations")
	}

	// make sure the database has the necessary tables
	mig := migrate.NewMigrator(db, migrationsToExec)
	err = mig.Init(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "Error initializing migrations")
	}

	// actually run the migrations
	_, err = mig.Migrate(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "Error executing migrations")
	}
	return db, err
}
