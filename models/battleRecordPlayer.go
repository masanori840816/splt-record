package models

import (
	"github.com/uptrace/bun"

	dto "github.com/splt-record/dto"
)

type BattleRecordPlayer struct {
	bun.BaseModel  `bun:"table:battle_record_players,alias:brp"`
	ID             int64 `bun:"id,pk,autoincrement" json:"id"`
	BattleRecordID int64 `bun:"battle_record_id,notnull,type:bigint" json:"battleRecordId"`
	WeaponID       int64 `bun:"weapon_id,notnull,type:bigint" json:"weaponId"`
	AliedPlayer    bool  `bun:"alied_player,notnull,type:boolean" json:"aliedPlayer"`

	BattleRecord *BattleRecord `bun:"rel:has-one,join:battle_record_id=id" json:"battleRecord"`
	Weapon       *Weapon       `bun:"rel:has-one,join:weapon_id=id" json:"weapon"`
}

func NewBattleRecordPlayer(player dto.BattleRecordPlayerForUpdate, recordID int64) *BattleRecordPlayer {
	return &BattleRecordPlayer{
		BattleRecordID: recordID,
		WeaponID:       player.WeaponID,
		AliedPlayer:    player.AliedPlayer,
	}
}
