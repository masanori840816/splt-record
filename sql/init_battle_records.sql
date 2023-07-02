CREATE TABLE battle_records
(id serial PRIMARY KEY,
battle_result_id bigint not null REFERENCES battle_results(id),
battle_date timestamp with time zone not null,
file_name varchar(256) not null
);

CREATE TABLE battle_record_players
(id serial PRIMARY KEY,
battle_record_id bigint not null REFERENCES battle_records(id),
weapon_id bigint not null REFERENCES weapons(id),
alied_player boolean not null
);

