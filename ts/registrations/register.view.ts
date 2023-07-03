import * as models from "../spltRecord.type";

export class RegisterView {
    private recodeDate: HTMLInputElement;
    private stageList: HTMLDataListElement;
    private stageInput: HTMLInputElement;
    private resultInput: HTMLSelectElement;
    private weaponList: HTMLDataListElement;
    private playerAlly1: HTMLInputElement;
    private playerAlly2: HTMLInputElement;
    private playerAlly3: HTMLInputElement;
    private playerAlly4: HTMLInputElement;
    private playerOpponent1: HTMLInputElement;
    private playerOpponent2: HTMLInputElement;
    private playerOpponent3: HTMLInputElement;
    private playerOpponent4: HTMLInputElement;
    constructor() {
        this.recodeDate = document.getElementById("register_record_date") as HTMLInputElement;
        this.stageList = document.getElementById("register_record_stage_list") as HTMLDataListElement;
        this.stageInput = document.getElementById("register_record_stage") as HTMLInputElement;
        this.resultInput = document.getElementById("register_record_result_input") as HTMLSelectElement;
        this.weaponList = document.getElementById("register_record_weapons_list") as HTMLDataListElement;
        this.playerAlly1 = document.getElementById("register_player_ally_1") as HTMLInputElement;
        this.playerAlly2 = document.getElementById("register_player_ally_2") as HTMLInputElement;
        this.playerAlly3 = document.getElementById("register_player_ally_3") as HTMLInputElement;
        this.playerAlly4 = document.getElementById("register_player_ally_4") as HTMLInputElement;
        this.playerOpponent1 = document.getElementById("register_player_opponent_1") as HTMLInputElement;
        this.playerOpponent2 = document.getElementById("register_player_opponent_2") as HTMLInputElement;
        this.playerOpponent3 = document.getElementById("register_player_opponent_3") as HTMLInputElement;
        this.playerOpponent4 = document.getElementById("register_player_opponent_4") as HTMLInputElement;

        console.log(this.recodeDate);
        console.log(this.stageInput);
        console.log(this.resultInput);
        console.log(this.playerAlly1);
        console.log(this.playerAlly2);
        console.log(this.playerAlly3);
        console.log(this.playerAlly4);
        console.log(this.playerOpponent1);        
        console.log(this.playerOpponent2);
        console.log(this.playerOpponent3);        
        console.log(this.playerOpponent4);
    }
    public setStageList(values: models.BattleStage[]) {
        for(const v of values) {
            const o = document.createElement("option");
            o.value = v.name;
            this.stageList.appendChild(o);
        }
    }
    public setWeaponList(values: models.Weapon[]) {
        for(const v of values) {
            const o = document.createElement("option");
            o.value = v.name;
            this.weaponList.appendChild(o);
        }
    }
}
