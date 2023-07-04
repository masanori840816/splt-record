package main

import (
	"context"

	db "github.com/splt-record/db"
	models "github.com/splt-record/models"
)

func GetAllWeapons(dbCtx *db.SpltRecordContext) ([]models.Weapon, error) {
	ctx := context.Background()
	return dbCtx.Weapon.GetAll(&ctx)
}
