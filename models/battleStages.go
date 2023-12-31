package models

import (
	"github.com/uptrace/bun"
)

type BattleStage struct {
	bun.BaseModel `bun:"table:battle_stages,alias:bst"`
	ID            int64  `bun:"id,pk,autoincrement" json:"id"`
	Name          string `bun:"name,notnull,type:varchar(128)" json:"name"`
}
