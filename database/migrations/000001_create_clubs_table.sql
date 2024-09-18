--liquibase formatted sql

-- changeset reza:1
CREATE TABLE IF NOT EXISTS clubs (
    id BIGSERIAL PRIMARY KEY,
    name varchar(64) NOT NULL UNIQUE,
    nationality varchar(3) NOT NULL,
    seed INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
-- rollback DROP TABLE IF EXISTS clubs;