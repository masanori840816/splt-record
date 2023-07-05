package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	db "github.com/splt-record/db"
	dto "github.com/splt-record/dto"
	files "github.com/splt-record/files"
)

func CreateRecord(w http.ResponseWriter, r *http.Request, dbCtx *db.SpltRecordContext) {
	w.Header().Set("Content-Type", "application/json")
	record, fileName, fileData, err := readRecords(r)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(400)
		return
	}
	ctx := context.Background()
	validateResult := dbCtx.Battles.ValidateForUpdate(&ctx, record)
	if !validateResult.Succeeded {
		resJSON, _ := json.Marshal(validateResult)
		w.Write(resJSON)
		return
	}
	err = files.SaveFile("record_images", fileName, fileData)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(400)
		return
	}
	result := dbCtx.Battles.Create(&ctx, record, fileName)
	resJSON, _ := json.Marshal(result)
	w.Write(resJSON)
}
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
func GetAllRules(w http.ResponseWriter, dbCtx *db.SpltRecordContext) {
	ctx := context.Background()
	results, err := dbCtx.Battles.GetAllRules(&ctx)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(502)
	}
	resJSON, _ := json.Marshal(results)
	w.Write(resJSON)
}
func readRecords(r *http.Request) (dto.BattleRecordForUpdate, string, []byte, error) {
	form, err := r.MultipartReader()
	if err != nil {
		return dto.BattleRecordForUpdate{}, "", nil, err
	}
	record := &dto.BattleRecordForUpdate{}
	fileName := ""
	var fileData []byte
	for {
		part, err := form.NextPart()
		if err == io.EOF {
			break
		}
		readData, err := io.ReadAll(part)
		if err != nil {
			log.Println(err.Error())
		}
		name := part.FormName()
		switch name {
		case "recordimage":
			fileData = readData
		case "filename":
			fileName = string(readData)
		case "record":
			err = json.Unmarshal(readData, &record)
			if err != nil {
				return dto.BattleRecordForUpdate{}, "", nil, err
			}
		}
	}
	return *record, fileName, fileData, nil
}
