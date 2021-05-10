package migration

import (
	"context"
	"sort"
	"time"

	infra "github.com/parpa/golang-gin-openapi/infrastructure"
)

type migrationInterface interface {
	GetId() int64
	Up(ctx context.Context, client *infra.Db) (bool, error)
	Down(ctx context.Context, client *infra.Db) (bool, error)
}

type migration struct {
	migrationInterface

	ctx    context.Context
	client *infra.Db
}

// var MigrationList migrationList

var migrationList = make(map[int64]migrationInterface)

func Up() (res bool, err error) {
	client := &infra.Db{}

	ctx := context.Background()
	m := &DbMigration{
		ctx:    ctx,
		client: client,
	}
	id, err := m.getMigrationId()
	if err != nil {
		return
	}
	// sort
	keys := make([]int64, 0, len(migrationList))
	for k := range migrationList {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	// run
	for _, key := range keys {
		if key <= id {
			continue
		}
		record := migrationList[key]
		res, err = record.Up(ctx, client)
		if err != nil || !res {
			return
		}
		err = m.setMigrationId(key)
		if err != nil {
			return
		}
	}

	return
}

func Down(client *infra.Db) (res bool, err error) {
	// todo

	return
}

type DbMigration struct {
	Id        int64     `json:"id"`        // migrationId
	UpdatedAt time.Time `json:"updatedAt"` // 更新日時

	ctx    context.Context
	client *infra.Db
}

func (m *DbMigration) getMigrationId() (int64, error) {
	// todo set m.Id

	return m.Id, nil
}

func (m *DbMigration) setMigrationId(id int64) error {
	m.Id = id
	m.UpdatedAt = time.Now()

	// todo set id

	return nil
}
