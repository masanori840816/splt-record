package models

import (
	"time"

	"github.com/uptrace/bun"

	dto "github.com/splt-record/dto"
)

type BattleRecord struct {
	bun.BaseModel  `bun:"table:battle_records,alias:brd"`
	ID             int64         `bun:"id,pk,autoincrement" json:"id"`
	BattleResultID int64         `bun:"battle_result_id,notnull,type:bigint" json:"battleResultId"`
	BattleStageID  int64         `bun:"battle_stage_id,notnull,type:bigint" json:"battleStageId"`
	BattleDate     time.Time     `bun:"battle_date,notnull,type:timestamp with time zone" json:"battleDate"`
	FileName       string        `bun:"file_name,notnull,type:varchar(256)" json:"fileName"`
	BattleResult   *BattleResult `bun:"rel:has-one,join:battle_result_id=id" json:"battleResult"`
	BattleStage    *BattleStage  `bun:"rel:has-one,join:battle_stage_id=id" json:"battleStage"`
}

func NewBattleRecord(record dto.BattleRecordForUpdate, fileName string) BattleRecord {
	return BattleRecord{
		BattleResultID: record.BattleResultID,
		BattleStageID:  record.BattleStageID,
		BattleDate:     record.BattleDate,
		FileName:       fileName,
	}
}
