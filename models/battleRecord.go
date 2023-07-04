package models

import (
	"time"

	"github.com/uptrace/bun"
)

type BattleRecord struct {
	bun.BaseModel  `bun:"table:battle_records,alias:brd"`
	ID             int64         `bun:"id,pk,autoincrement"`
	BattleResultID int64         `bun:"battle_result_id,notnull,type:bigint"`
	BattleStageID  int64         `bun:"battle_stage_id,notnull,type:bigint"`
	BattleDate     time.Time     `bun:"battle_date,notnull,type:timestamp with time zone"`
	FileName       string        `bun:"file_name,notnull,type:varchar(256)"`
	BattleResult   *BattleResult `bun:"rel:has-one,join:battle_result_id=id"`
	BattleStage    *BattleStage  `bun:"rel:has-one,join:battle_stage_id=id"`
}
