-- +migrate Up
CREATE TABLE shop(id text primary key, name text not null, introduction text not null, location text not null, level int not null, score float not null);

-- +migrate Down
DROP TABLE shop;