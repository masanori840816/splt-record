import * as models from "./spltRecord.type";
import * as web from "./registrations/registers";
import { RegisterView } from "./registrations/register.view";

let view: RegisterView;
let stages: models.BattleStage[];
let weapons: models.Weapon[];
window.RegisterPage = {    
    async init(baseUrl: string) {
        view = new RegisterView();
        stages = await web.getAllStages(baseUrl);
        weapons = await web.getAllWeapons(baseUrl);
        view.setStageList(stages);
        view.setWeaponList(weapons);
        view.setResultList(await web.getAllResults(baseUrl));
    },
    save(baseUrl: string){

        console.log(baseUrl);
        const newRec = view.generateRecord(weapons);
        console.log(newRec);
        
    },
}

