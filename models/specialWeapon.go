package models

import (
	"github.com/uptrace/bun"
)

type SpecialWeapon struct {
	bun.BaseModel `bun:"table:special_weapons,alias:spw"`
	ID            int64  `bun:"id,pk,autoincrement" json:"id"`
	Name          string `bun:"name,notnull,type:varchar(128)" json:"name"`
	Power         int32  `bun:"power,notnull,default:1,type:integer" json:"power"`
}
