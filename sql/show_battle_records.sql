SELECT rcd.id as "BattleID", ply.id AS "PlayerID", ply.alied_player,
wpn.name as "Weapon", stg.name as "Stage", rul.name as "Rule",
rlt.name as "Result", rcd.battle_date as "Date", rcd.file_name as "File"  
FROM battle_record_players ply
INNER JOIN battle_records rcd ON ply.battle_record_id = rcd.id
INNER JOIN weapons wpn ON ply.weapon_id = wpn.id
INNER JOIN battle_results rlt ON rcd.battle_result_id = rlt.id
INNER JOIN battle_stages stg ON rcd.battle_stage_id = stg.id
INNER JOIN battle_rules rul ON rcd.battle_rule_id = rul.id
ORDER BY rcd.id, ply.id;