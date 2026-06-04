-- Baseline schema. Mirrors the GORM models in internal/domain so that
-- AutoMigrate (dev) and these migrations (prod) converge on the same shape:
-- identical table, column, index, and constraint names => no drift.

CREATE TABLE IF NOT EXISTS accounts (
    id            BIGSERIAL PRIMARY KEY,
    email         TEXT        NOT NULL,
    password_hash TEXT        NOT NULL,
    full_name     TEXT        NOT NULL,
    role          TEXT        NOT NULL DEFAULT 'owner',
    created_at    TIMESTAMPTZ,
    updated_at    TIMESTAMPTZ
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_accounts_email ON accounts (email);

CREATE TABLE IF NOT EXISTS clinics (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT        NOT NULL,
    address     TEXT,
    phone       TEXT,
    email       TEXT,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ
);
CREATE INDEX IF NOT EXISTS idx_clinics_name ON clinics (name);

CREATE TABLE IF NOT EXISTS doctors (
    id          BIGSERIAL PRIMARY KEY,
    clinic_id   BIGINT      NOT NULL,
    full_name   TEXT        NOT NULL,
    specialty   TEXT,
    license_no  TEXT,
    phone       TEXT,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    CONSTRAINT fk_doctors_clinic FOREIGN KEY (clinic_id) REFERENCES clinics (id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS idx_doctors_clinic_id ON doctors (clinic_id);
CREATE INDEX IF NOT EXISTS idx_doctors_specialty ON doctors (specialty);

CREATE TABLE IF NOT EXISTS pets (
    id          BIGSERIAL PRIMARY KEY,
    owner_id    BIGINT,
    name        TEXT        NOT NULL,
    species     TEXT        NOT NULL,
    breed       TEXT,
    age         BIGINT,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    CONSTRAINT fk_pets_owner FOREIGN KEY (owner_id) REFERENCES accounts (id) ON DELETE SET NULL
);
CREATE INDEX IF NOT EXISTS idx_pets_owner_id ON pets (owner_id);
