package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	dto "github.com/splt-record/dto"
	models "github.com/splt-record/models"
	"github.com/uptrace/bun"
)

type Battles struct {
	db *bun.DB
}

func NewBattles(database *bun.DB) *Battles {
	return &Battles{
		db: database,
	}
}
func (b Battles) Create(ctx *context.Context, record dto.BattleRecordForUpdate, fileName string) dto.ActionResult {
	tx, err := b.db.BeginTx(*ctx, &sql.TxOptions{})
	if err != nil {
		log.Println(err.Error())
		return dto.GetFailedResult("CREATE RECORD ERROR")
	}
	newRecord := models.NewBattleRecord(record, fileName)
	_, err = tx.NewInsert().Model(newRecord).Exec(*ctx)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return dto.GetFailedResult("CREATE RECORD ERROR")
	}
	for _, p := range record.Players {
		newPlayer := models.NewBattleRecordPlayer(p, newRecord.ID)
		_, err = tx.NewInsert().Model(newPlayer).Exec(*ctx)
		if err != nil {
			log.Println(err.Error())
			tx.Rollback()
			return dto.GetFailedResult("CREATE RECORD ERROR")
		}
	}

	tx.Commit()
	return dto.GetSucceededResult()
}
func (b Battles) GetAllResults(ctx *context.Context) ([]models.BattleResult, error) {
	results := make([]models.BattleResult, 0)
	err := b.db.NewSelect().
		Model(&results).
		Scan(*ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (b Battles) GetAllStages(ctx *context.Context) ([]models.BattleStage, error) {
	results := make([]models.BattleStage, 0)
	err := b.db.NewSelect().
		Model(&results).
		Scan(*ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (b Battles) GetAllRules(ctx *context.Context) ([]models.BattleRule, error) {
	results := make([]models.BattleRule, 0)
	err := b.db.NewSelect().
		Model(&results).
		Scan(*ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// Validation for creating and updating records.
func (b Battles) ValidateForUpdate(ctx *context.Context, record dto.BattleRecordForUpdate) dto.ActionResult {
	if record.ID > 0 {
		exists, err := b.db.NewSelect().Model(new(models.BattleRecord)).
			Where("id=?", record.ID).Exists(*ctx)
		if err != nil {
			log.Println(err.Error())
			return dto.GetFailedResult("FAILED VALIDATION")
		}
		if !exists {
			return dto.GetFailedResult(fmt.Sprintf("INVALID BATTLE ID:%d", record.ID))
		}
	}
	exists, err := b.db.NewSelect().Model(new(models.BattleResult)).
		Where("id=?", record.BattleResultID).Exists(*ctx)
	if err != nil {
		log.Println(err.Error())
		return dto.GetFailedResult("FAILED VALIDATION")
	}
	if !exists {
		return dto.GetFailedResult(fmt.Sprintf("INVALID BATTLE RESULT ID:%d", record.BattleResultID))
	}
	exists, err = b.db.NewSelect().Model(new(models.BattleStage)).
		Where("id=?", record.BattleStageID).Exists(*ctx)
	if err != nil {
		log.Println(err.Error())
		return dto.GetFailedResult("FAILED VALIDATION")
	}
	if !exists {
		return dto.GetFailedResult(fmt.Sprintf("INVALID BATTLE STAGE ID:%d", record.BattleStageID))
	}
	exists, err = b.db.NewSelect().Model(new(models.BattleRule)).
		Where("id=?", record.BattleRuleID).Exists(*ctx)
	if err != nil {
		log.Println(err.Error())
		return dto.GetFailedResult("FAILED VALIDATION")
	}
	if !exists {
		return dto.GetFailedResult(fmt.Sprintf("INVALID BATTLE RULE ID:%d", record.BattleRuleID))
	}
	if record.BattleDate.Year() < 2023 {
		return dto.GetFailedResult(fmt.Sprintf("INVALID BATTLE DATE:%s", record.BattleDate.Format("2006-01-02 15:04")))
	}
	playerIDs := ""
	weaponIDs := ""
	for _, player := range record.Players {
		// Skip validating new data
		if player.ID > 0 {
			playerIDs = fmt.Sprintf("%s, %d", playerIDs, player.ID)
		}
		// All weapon IDs must be existed.
		weaponIDs = fmt.Sprintf("%s, %d", weaponIDs, player.WeaponID)
	}
	playerIDs = strings.TrimPrefix(playerIDs, ",")
	weaponIDs = strings.TrimPrefix(weaponIDs, ",")
	// Extract IDs that do not exist in the database
	if len(playerIDs) > 0 {
		invalidPlayerIDs := make([]int, 0)
		err = b.db.NewRaw(
			fmt.Sprintf(`SELECT t1.keyword FROM (SELECT UNNEST(ARRAY[%s]) as keyword) as t1
			LEFT JOIN battle_record_players t2 ON t1.keyword =t2.ID	WHERE t2.ID is null
			`, playerIDs)).Scan(*ctx, &invalidPlayerIDs)
		if err != nil {
			log.Println(err.Error())
			return dto.GetFailedResult("FAILED VALIDATION")
		}
		if len(invalidPlayerIDs) > 0 {
			invalidIDText := ""
			for _, id := range invalidPlayerIDs {
				invalidIDText = fmt.Sprintf("%s, %d", invalidIDText, id)
			}

			return dto.GetFailedResult(fmt.Sprintf("INVALID PLAYER IDS%s", invalidIDText))
		}
	}
	invalidWeaponIDs := make([]int, 0)
	err = b.db.NewRaw(
		fmt.Sprintf(`SELECT t1.keyword FROM (SELECT UNNEST(ARRAY[%s]) as keyword) as t1
		LEFT JOIN weapons t2 ON t1.keyword =t2.ID	WHERE t2.ID is null
        `, weaponIDs)).Scan(*ctx, &invalidWeaponIDs)
	if err != nil {
		log.Println(err.Error())
		return dto.GetFailedResult("FAILED VALIDATION")
	}
	if len(invalidWeaponIDs) > 0 {
		invalidIDText := ""
		for _, id := range invalidWeaponIDs {
			invalidIDText = fmt.Sprintf("%s, %d", invalidIDText, id)
		}
		return dto.GetFailedResult(fmt.Sprintf("INVALID WEAPON IDS%s", invalidIDText))
	}
	return dto.GetSucceededResult()
}
