import * as models from "./spltRecord.type";
import * as web from "./registrations/registers";
import { RegisterView } from "./registrations/register.view";
import { hasAnyTexts } from "./texts/hasAnyTexts";

let view: RegisterView;
let weapons: models.Weapon[];
window.RegisterPage = {    
    async init(baseUrl: string) {
        view = new RegisterView();
        weapons = await web.getAllWeapons(baseUrl);
        view.setStageList(await web.getAllStages(baseUrl));
        view.setWeaponList(weapons);
        view.setRuleList(await web.getAllRules(baseUrl));
        view.setResultList(await web.getAllResults(baseUrl));
    },
    async save(baseUrl: string){
        const record = view.generateRecord(weapons);
        const file = view.getFile();
        if(validateRecord(record, file)) {
            const result = await web.createRecord(baseUrl, record, file!);
            if(result.succeeded) {
                alert("登録完了");
            } else {
                alert(result.errorMessage);
            }
        } else {
            alert("入力エラー");
        }
        
    },
}
function validateRecord(record: models.BattleRecordForUpdate, file: models.UploadFile|null): boolean {
    if(record.battleStageId <= 0 ||
        record.battleDate.getFullYear() < 2023 ||
        record.players.length < 8) {
        return false;
    }
    for(const p of record.players) {
        if(p.weaponId <= 0) {
            return false;
        }
    }
    if(file == null ||
        hasAnyTexts(file.name) === false ||
        file.fileData.size <= 0) {
        return false;
    }
    return true;
}

