--liquibase formatted sql
--changeset om-nikam:1
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--rollback DROP EXTENSION "uuid-ossp";

--changeset om-nikam:2
create table user
(
    "id"           UUID primary key,
    "first_name"   varchar        not null,
    "last_name"    varchar,
    "email"        varchar unique not null,
    "password"     varchar        not null,
    "verified"     bool,
);
--rollback DROP TABLE user;

create table course
(
    "id"          UUID primary key,
    "user_id"     UUID not null,
    "title"       varchar,
    "price"       float,
    "description" varchar,
    foreign key (user_id) references user (id)
        on delete cascade
);
--rollback DROP TABLE course;