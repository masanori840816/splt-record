package dto

import "time"

type BattleRecordForUpdate struct {
	ID             int64                         `json:"id"`
	BattleResultID int64                         `json:"battleResultId"`
	BattleStageID  int64                         `json:"battleStageId"`
	BattleDate     time.Time                     `json:"battleDate"`
	Players        []BattleRecordPlayerForUpdate `json:"players"`
}

type BattleRecordPlayerForUpdate struct {
	ID          int64 `json:"id"`
	WeaponID    int64 `json:"weaponId"`
	AliedPlayer bool  `json:"aliedPlayer"`
}
