package models

import (
	"github.com/uptrace/bun"
)

type BattleResult struct {
	bun.BaseModel `bun:"table:battle_results,alias:brs"`
	ID            int64  `bun:"id,pk,autoincrement" json:"id"`
	Name          string `bun:"name,notnull,type:varchar(32)" json:"name"`
}
