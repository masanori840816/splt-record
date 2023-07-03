import * as models from "./spltRecord.type";
import { RegisterView } from "./registrations/register.view";

let view: RegisterView;
window.RegisterPage = {    
    init(baseUrl: string) {
        view = new RegisterView();
        const sampleStages: models.BattleStage[] = [
            {
                id: 1,
                name: "サンプル１"
            },
            {
                id: 2,
                name: "サンプル2"
            },
        ];
        const sampleWeapons: models.Weapon[] = [
            {
                id: 1,
                name: "武器１",
            },
            {
                id: 2,
                name: "武器２",
            }
        ];

        console.log(baseUrl);
        view.setStageList(sampleStages);
        view.setWeaponList(sampleWeapons);
        
    },
    save(baseUrl: string){

        console.log(baseUrl);
    },
}

