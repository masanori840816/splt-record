import * as models from "../spltRecord.type";

export async function getAllWeapons(baseUrl: string): Promise<models.Weapon[]> {
    try {
        const response = await fetch(`${baseUrl}weapons/all`, {
            "method": "GET"
        });
        const j = await response.json();
        return JSON.parse(JSON.stringify(j))
    }catch(err) {
        console.error(err);
        return [];
    }
}

export async function getAllStages(baseUrl: string): Promise<models.BattleStage[]> {
    try {
        const response = await fetch(`${baseUrl}battle-stages/all`, {
            "method": "GET"
        });
        const j = await response.json();
        return JSON.parse(JSON.stringify(j))
    }catch(err) {
        console.error(err);
        return [];
    }
}

export async function getAllResults(baseUrl: string): Promise<models.BattleResult[]> {
    try {
        const response = await fetch(`${baseUrl}battle-results/all`, {
            "method": "GET"
        });
        const j = await response.json();
        return JSON.parse(JSON.stringify(j))
    }catch(err) {
        console.error(err);
        return [];
    }
}
export async function getAllRules(baseUrl: string): Promise<models.BattleRule[]> {
    try {
        const response = await fetch(`${baseUrl}battle-rules/all`, {
            "method": "GET"
        });
        const j = await response.json();
        return JSON.parse(JSON.stringify(j))
    }catch(err) {
        console.error(err);
        return [];
    }
}
export async function createRecord(baseUrl: string, record: models.BattleRecordForUpdate, file: models.UploadFile): Promise<models.ActionResult> {
    console.log(`${baseUrl} ${file}`);
    console.log(record);
    
    
    return {
        succeeded: false,
        errorMessage: "not implemented"
    };
}