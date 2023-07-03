package db

import "github.com/uptrace/bun"

type Weapons struct {
	db *bun.DB
}

func NewWeapons(database *bun.DB) *Weapons {
	return &Weapons{
		db: database,
	}
}
