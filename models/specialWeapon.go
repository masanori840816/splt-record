package models

import (
	"github.com/uptrace/bun"
)

type SpecialWeapon struct {
	bun.BaseModel `bun:"table:special_weapons,alias:spw"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull,type:varchar(128)"`
	Power         int32  `bun:"power,notnull,type:integer,default=1"`
}
