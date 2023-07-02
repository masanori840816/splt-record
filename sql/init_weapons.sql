CREATE TABLE weapons
(id serial PRIMARY KEY,
weapon_type_id bigint not null REFERENCES weapon_types(id),
sub_weapon_id bigint not null REFERENCES sub_weapons(id),
special_weapon_id bigint not null REFERENCES special_weapons(id),
name varchar(128) not null,
range float not null default 1
);


INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 1, 5, 'わかばシューター', 1.5);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 2, 1, 'ヒーローシューターレプリカ', 2.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 2, 1, 'スプラシューター', 2.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 6, 6, 'プロモデラーMG', 2.1);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 2, 2, 'N-ZAP85', 2.3);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 1, 8, 'スプラシューターコラボ', 2.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 8, 14, 'プロモデラーRG', 2.1);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 5, 17, 'N-ZAP89', 2.3);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 10, 9, 'もみじシューター', 1.5);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 13, 10, 'スペースシューター', 3.3);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 12, 3, 'プライムシューター', 3.5);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 3, 12, 'ボールドマーカー', 1.5);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 9, 13, 'スペースシューターコラボ', 3.3);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 2, 14, 'プライムシューターコラボ', 3.5);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 11, 10, 'ボールドマーカーネオ', 1.5);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 7, 10, '52ガロン', 2.7);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 3, 3, 'L3リールガン', 3.0);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 13, 2, 'H3リールガン', 3.3);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 12, 4, 'ジェットスイーパー', 4.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 4, 12, 'L3リールガンD', 3.0);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 7, 5, 'H3リールガンD', 3.3);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 14, 11, 'ジェットスイーパーカスタム', 4.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 4, 3, 'シャープマーカー', 2.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 8, 4, '96ガロン', 3.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 7, 1, 'ボトルガイザー', 3.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 2, 8, 'シャープマーカーネオ', 2.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(1, 7, 16, '96ガロンデコ', 3.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(2, 2, 3, 'スプラマニューバー', 2.5);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(2, 1, 9, 'デュアルスイーパー', 3.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(2, 11, 2, 'スパッタリー', 1.5);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(2, 5, 6, 'クアッドホッパーブラック', 2.8);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(2, 7, 14, 'ケルビン525', 3.1);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(2, 11, 17, 'デュアルスイーパーカスタム', 3.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(2, 10, 6, 'スパッタリー・ヒュー', 1.5);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(2, 8, 7, 'クアッドホッパーホワイト', 2.8);

INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(3, 5, 5, 'ホットブラスター', 2.9);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(3, 9, 8, 'ラピッドブラスター', 3.8);

INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(3, 10, 13, 'ラピッドブラスターデコ', 3.8);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(3, 2, 9, 'ロングブラスター', 3.4);

INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(3, 1, 7, 'ノヴァブラスター', 2.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(3, 6, 12, 'ノヴァブラスターネオ', 2.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(3, 8, 6, 'S-BLAST92', 2.8);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(3, 1, 1, 'クラッシュブラスター', 2.7);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(3, 3, 17, 'クラッシュブラスターネオ', 2.7);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(3, 14, 4, 'Rブラスターエリート', 4.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(3, 12, 10, 'Rブラスターエリートデコ', 4.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(4, 2, 7, 'ホクサイ', 1.1);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(4, 1, 10, 'パブロ', 1.1);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(4, 3, 9, 'フィンセント', 1.4);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(4, 9, 12, 'パブロ・ヒュー', 1.1);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(5, 3, 5, 'スプラローラー', 2.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(5, 5, 7, 'カーボンローラー', 2.5);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(5, 8, 2, 'ダイナモローラー', 3.0);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(5, 7, 4, 'ワイドローラー', 2.8);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(5, 12, 11, 'ワイドローラーコラボ', 2.8);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(5, 9, 15, 'ヴァリアブルローラー', 3.0);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(5, 11, 16, 'スプラローラーコラボ', 2.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(5, 4, 1, 'カーボンローラーデコ', 2.5);

INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(6, 1, 8, 'バケットスロッシャー', 3.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(6, 14, 13, 'ヒッセン', 2.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(6, 6, 14, 'スクリュースロッシャー', 3.0);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(6, 12, 7, 'バケットスロッシャーデコ', 3.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(6, 6, 2, 'ヒッセン・ヒュー', 2.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(6, 8, 11, 'オーバーフロッシャー', 5.0);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(6, 13, 11, 'エクスプロッシャー', 4.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(8, 1, 4, 'スプラチャージャー', 5.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(8, 13, 5, 'スクイックリンα', 3.9);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(8, 7, 8, 'スプラチャージャーコラボ', 5.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(8, 1, 4, 'スプラスコープ', 5.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(8, 7, 8, 'スプラスコープコラボ', 5.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(8, 8, 2, 'R-PEN/5H', 5.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(8, 9, 9, 'リッター4K', 6.2);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(8, 9, 9, '4Kスコープ', 6.6);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(8, 5, 10, '14式竹筒銃・甲', 4.3);
INSERT INTO weapons
(weapon_type_id, sub_weapon_id,
special_weapon_id, name, range)
VALUES
(8, 10, 15, 'ソイチューバー', 4.2);
