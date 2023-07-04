package models

import (
	"github.com/uptrace/bun"
)

type Weapon struct {
	bun.BaseModel   `bun:"table:weapons,alias:wpn"`
	ID              int64   `bun:"id,pk,autoincrement" json:"id"`
	WeaponTypeID    int64   `bun:"weapon_type_id,notnull,type:bigint" json:"weaponTypeId"`
	SubWeaponID     int64   `bun:"sub_weapon_id,notnull,type:bigint" json:"subWeaponId"`
	SpecialWeaponID int64   `bun:"special_weapon_id,notnull,type:bigint" json:"specialWeaponId"`
	Name            string  `bun:"name,notnull,type:varchar(128)" json:"name"`
	Range           float32 `bun:"range,notnull,type:float,default:1" json:"range"`

	WeaponType    *WeaponType    `bun:"rel:has-one,join:weapon_type_id=id" json:"weaponType"`
	SubWeapon     *SubWeapon     `bun:"rel:has-one,join:sub_weapon_id=id" json:"subWeapon"`
	SpecialWeapon *SpecialWeapon `bun:"rel:has-one,join:special_weapon_id=id" json:"specialWeapon"`
}
