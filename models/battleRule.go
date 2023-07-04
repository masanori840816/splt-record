package models

import (
	"github.com/uptrace/bun"
)

type BattleRule struct {
	bun.BaseModel `bun:"table:battle_rules,alias:brl"`
	ID            int64  `bun:"id,pk,type:integer" json:"id"`
	Name          string `bun:"name,notnull,type:varchar(64)" json:"name"`
}
