-- users
CREATE TABLE IF NOT EXISTS users (
    id            BIGSERIAL PRIMARY KEY,
    email         TEXT        NOT NULL UNIQUE,
    password_hash TEXT        NOT NULL,
    full_name     TEXT        NOT NULL,
    phone         TEXT        NOT NULL DEFAULT '',
    role          TEXT        NOT NULL DEFAULT 'owner',
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);

-- clinics
CREATE TABLE IF NOT EXISTS clinics (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT        NOT NULL,
    address     TEXT        NOT NULL DEFAULT '',
    phone       TEXT        NOT NULL DEFAULT '',
    email       TEXT        NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_clinics_name ON clinics(name);

-- doctors
CREATE TABLE IF NOT EXISTS doctors (
    id          BIGSERIAL PRIMARY KEY,
    clinic_id   BIGINT      NOT NULL REFERENCES clinics(id) ON DELETE CASCADE,
    full_name   TEXT        NOT NULL,
    specialty   TEXT        NOT NULL DEFAULT '',
    license_no  TEXT        NOT NULL DEFAULT '',
    phone       TEXT        NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_doctors_clinic_id ON doctors(clinic_id);
CREATE INDEX IF NOT EXISTS idx_doctors_specialty ON doctors(specialty);

-- pets
CREATE TABLE IF NOT EXISTS pets (
    id          BIGSERIAL PRIMARY KEY,
    owner_id    BIGINT      REFERENCES users(id) ON DELETE SET NULL,
    name        TEXT        NOT NULL,
    species     TEXT        NOT NULL,
    breed       TEXT        NOT NULL DEFAULT '',
    age         INTEGER     NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_pets_owner_id ON pets(owner_id);
CREATE INDEX IF NOT EXISTS idx_pets_species ON pets(species);
