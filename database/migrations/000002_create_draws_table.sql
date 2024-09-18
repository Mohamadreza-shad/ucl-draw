--liquibase formatted sql

-- changeset reza:2
CREATE TABLE IF NOT EXISTS draws (
    id BIGSERIAL PRIMARY KEY,
    host varchar(3) NOT NULL,
    guest varchar(3) NOT NULL
);
-- rollback DROP TABLE IF EXISTS draws;