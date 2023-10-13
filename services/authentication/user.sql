create table member (
    id serial primary key,
    email varchar(100) not null unique,
    username varchar(50) not null unique,
    password varchar not null,
    firstname varchar(30) not null,
    lastname varchar(30) not null,
    birthdate date not null,
    city varchar(40) not null,
    country varchar(40) not null
);