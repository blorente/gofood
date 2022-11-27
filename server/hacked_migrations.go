package main

//Copied from migrations.go
import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/migrations/logs"
	"github.com/pocketbase/pocketbase/tools/migrate"
)

type migrationsConnection struct {
	DB             *dbx.DB
	MigrationsList migrate.MigrationsList
}

func migrationsConnectionsMap(app core.App) map[string]migrationsConnection {
	return map[string]migrationsConnection{
		"db": {
			DB:             app.DB(),
			MigrationsList: migrations.AppMigrations,
		},
		"logs": {
			DB:             app.LogsDB(),
			MigrationsList: logs.LogsMigrations,
		},
	}
}
func runMigrations(app core.App) error {
	connections := migrationsConnectionsMap(app)

	for _, c := range connections {
		runner, err := migrate.NewRunner(c.DB, c.MigrationsList)
		if err != nil {
			return err
		}

		if _, err := runner.Up(); err != nil {
			return err
		}
	}

	return nil
}
