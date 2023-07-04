package models

import (
	"github.com/uptrace/bun"
)

type BattleRecordPlayer struct {
	bun.BaseModel  `bun:"table:battle_record_players,alias:brp"`
	ID             int64 `bun:"id,pk,autoincrement"`
	BattleRecordID int64 `bun:"battle_record_id,notnull,type:bigint"`
	WeaponID       int64 `bun:"weapon_id,notnull,type:bigint"`
	AliedPlayer    bool  `bun:"alied_player,notnull,type:boolean"`

	BattleRecord *BattleRecord `bun:"rel:has-one,join:battle_record_id=id"`
	Weapon       *Weapon       `bun:"rel:has-one,join:weapon_id=id"`
}
