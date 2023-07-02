CREATE TABLE battle_results
(id integer PRIMARY KEY,
name varchar(32) not null);

INSERT INTO battle_results
(id, name)
Values
(1, 'win');
INSERT INTO battle_results
(id, name)
Values
(2, 'lose');
INSERT INTO battle_results
(id, name)
Values
(3, 'draw');