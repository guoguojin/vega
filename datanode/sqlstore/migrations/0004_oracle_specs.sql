-- +goose Up
ALTER TABLE oracle_specs DROP COLUMN IF EXISTS signers;
ALTER TABLE oracle_specs DROP COLUMN IF EXISTS filters;
ALTER TABLE oracle_specs ADD COLUMN data JSONB NOT NULL;