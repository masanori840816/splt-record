package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type SpltRecordContext struct {
	db     *bun.DB
	Weapon *Weapons
}

func NewSpltRecordContext(connectionString string) *SpltRecordContext {
	result := SpltRecordContext{}
	result.db = bun.NewDB(
		sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString))),
		pgdialect.New(),
	)
	result.Weapon = NewWeapons(result.db)
	return &result
}
