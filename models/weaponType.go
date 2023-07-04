package models

import (
	"github.com/uptrace/bun"
)

type WeaponType struct {
	bun.BaseModel `bun:"table:weapon_types,alias:wty"`
	ID            int64  `bun:"id,pk,type:integer" json:"id"`
	Name          string `bun:"name,notnull,type:varchar(64)" json:"name"`
}
