package db

import (
	"context"

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
