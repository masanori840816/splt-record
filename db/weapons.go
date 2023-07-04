package db

import (
	"context"

	models "github.com/splt-record/models"
	"github.com/uptrace/bun"
)

type Weapons struct {
	db *bun.DB
}

func NewWeapons(database *bun.DB) *Weapons {
	return &Weapons{
		db: database,
	}
}
func (w Weapons) GetAll(ctx *context.Context) ([]models.Weapon, error) {
	results := make([]models.Weapon, 0)
	err := w.db.NewSelect().
		Model(&results).
		Relation("WeaponType").
		Relation("SubWeapon").
		Relation("SpecialWeapon").
		Scan(*ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}
