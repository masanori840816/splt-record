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
    try{
        const formData = new FormData();
        formData.append("recordimage", file.fileData);
        formData.append("filename", file.name);
        formData.append("record", JSON.stringify(record));
        const response = await fetch(`${baseUrl}records`,{
            method: "POST",
            body: formData
        });
        const jsn = await response.json();
        return JSON.parse(JSON.stringify(jsn));
    }catch(err) {
        console.error(err);   
    }
    return {
        succeeded: false,
        errorMessage: "failed"
    };
}