export type BattleStage = {
    id: number,
    name: string
}
export type Weapon = {
    id: number,
    name: string,
}

export type BattleResult = {
	id: number,
    name: string,
}
export type BattleRecordForUpdate = {
	id: number,
	battleResultId: number,
	battleStageId: number,
	battleDate: Date,
    players: BattleRecordPlayerForUpdate[]
}

export type BattleRecordPlayerForUpdate = {
	id: number,
	weaponId: number,
	aliedPlayer: boolean,
}
export type UploadFile = {
	name: string,
	fileData: Blob
}
export type ActionResult = {
	succeeded: boolean,
	errorMessage: string
}