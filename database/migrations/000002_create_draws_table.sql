--liquibase formatted sql

-- changeset reza:2
CREATE TABLE IF NOT EXISTS draws (
    id BIGSERIAL PRIMARY KEY,
    host_id INT NOT NULL,
    guest_id INT NOT NULL
);
-- rollback DROP TABLE IF EXISTS draws;