package models

import (
	"github.com/uptrace/bun"
)

type Weapon struct {
	bun.BaseModel   `bun:"table:weapons,alias:wpn"`
	ID              int64   `bun:"id,pk,autoincrement"`
	WeaponTypeID    int64   `bun:"weapon_type_id,notnull,type:bigint"`
	SubWeaponID     int64   `bun:"sub_weapon_id,notnull,type:bigint"`
	SpecialWeaponID int64   `bun:"special_weapon_id,notnull,type:bigint"`
	Name            string  `bun:"name,notnull,type:varchar(128)"`
	Range           float32 `bun:"range,notnull,type:float,default:1"`

	WeaponType    *WeaponType    `bun:"rel:has-one,join:weapon_type_id=id"`
	SubWeapon     *SubWeapon     `bun:"rel:has-one,join:sub_weapon_id=id"`
	SpecialWeapon *SpecialWeapon `bun:"rel:has-one,join:special_weapon_id=id"`
}
