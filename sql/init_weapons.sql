CREATE TABLE weapons
(id serial PRIMARY KEY,
weapon_type_id bigint not null REFERENCES weapon_types(id),
sub_weapon_id bigint not null REFERENCES sub_weapons(id),
special_weapon_id bigint not null REFERENCES special_weapons(id),
name varchar(128) not null,
range float not null default 1
);
