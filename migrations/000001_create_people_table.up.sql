CREATE TABLE people (
    id          serial     PRIMARY KEY,
    name        varchar(50),
    surname     varchar(50),
    patronymic  varchar(50),
    age         int,
    gender      varchar(10),
    country     varchar(50)
);