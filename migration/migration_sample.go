package migration

import (
	"context"

	infra "github.com/parpa/golang-gin-openapi/infrastructure"
)

type migration_1619844594826 struct {
	*migration
}

func init() {
	// register
	migrationList[1234567890123] = &migration_1619844594826{
		migration: &migration{},
	}
}

func (m *migration_1619844594826) Up(ctx context.Context, client *infra.Db) (res bool, err error) {
	m.client = client
	m.ctx = ctx

	// todo update

	return true, nil
}

func (m *migration_1619844594826) Down(ctx context.Context, client *infra.Db) (res bool, err error) {
	m.client = client
	m.ctx = ctx

	// todo

	return
}
