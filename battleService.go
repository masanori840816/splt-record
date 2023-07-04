package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	db "github.com/splt-record/db"
)

func GetAllBattleResult(w http.ResponseWriter, dbCtx *db.SpltRecordContext) {
	ctx := context.Background()
	results, err := dbCtx.Battles.GetAllResults(&ctx)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(502)
	}
	resJSON, _ := json.Marshal(results)
	w.Write(resJSON)
}
func GetAllStages(w http.ResponseWriter, dbCtx *db.SpltRecordContext) {
	ctx := context.Background()
	results, err := dbCtx.Battles.GetAllStages(&ctx)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(502)
	}
	resJSON, _ := json.Marshal(results)
	w.Write(resJSON)
}
