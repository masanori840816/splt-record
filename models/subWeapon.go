package models

import (
	"github.com/uptrace/bun"
)

type SubWeapon struct {
	bun.BaseModel `bun:"table:sub_weapons,alias:swp"`
	ID            int64  `bun:"id,pk,autoincrement" json:"id"`
	Name          string `bun:"name,notnull,type:varchar(128)" json:"name"`
}
