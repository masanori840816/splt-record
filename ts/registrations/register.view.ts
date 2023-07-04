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
    public setResultList(values: models.BattleResult[]) {
        for(const v of values) {
            const o = document.createElement("option");
            o.value = v.id.toString();
            o.textContent = v.name;
            this.resultInput.appendChild(o);
        }
    }
    public generateRecord(values: models.Weapon[]): models.BattleRecordForUpdate {
        return {
            id: -1,
            battleDate: new Date(),
            battleStageId: -1,
            battleResultId: parseInt(this.resultInput.options[this.resultInput.selectedIndex]!.value),
            players: this.generateRecordPlayers(values),
        };
    }
    private generateRecordPlayers(values: models.Weapon[]): models.BattleRecordPlayerForUpdate[] {
        return [
            {
                id: -1,
                weaponId: this.getWeaponId(this.playerAlly1.value.trim(), values),
                aliedPlayer: true,
            },
            {
                id: -1,
                weaponId: this.getWeaponId(this.playerAlly2.value.trim(), values),
                aliedPlayer: true,
            },
            {
                id: -1,
                weaponId: this.getWeaponId(this.playerAlly3.value.trim(), values),
                aliedPlayer: true,
            },
            {
                id: -1,
                weaponId: this.getWeaponId(this.playerAlly4.value.trim(), values),
                aliedPlayer: true,
            },
            {
                id: -1,
                weaponId: this.getWeaponId(this.playerOpponent1.value.trim(), values),
                aliedPlayer: false,
            },
            {
                id: -1,
                weaponId: this.getWeaponId(this.playerOpponent2.value.trim(), values),
                aliedPlayer: false,
            },
            {
                id: -1,
                weaponId: this.getWeaponId(this.playerOpponent3.value.trim(), values),
                aliedPlayer: false,
            },
            {
                id: -1,
                weaponId: this.getWeaponId(this.playerOpponent4.value.trim(), values),
                aliedPlayer: false,
            },
        ];
    }
    private getWeaponId(name: string, values: models.Weapon[]): number {
        for(const v of values) {
            if(v.name === name) {
                return v.id;
            }
        }
        return -1;
    }
}
