CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "role" AS ENUM (
  'user',
  'admin',
  'employee'
);

CREATE TYPE "id_card_type" AS ENUM (
  'KTP',
  'SIM',
  'Passport'
);

CREATE TYPE "transaction_type" AS ENUM (
  'topup',
  'withdrawal',
  'transfer',
  'payment'
);

CREATE TYPE "customer_enum" AS ENUM (
  'active',
  'inactive',
  'pending',
  'blocked'
);

CREATE TABLE "m_customer" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "role" role NOT NULL DEFAULT 'user',
  "id_card_type" id_card_type NOT NULL DEFAULT 'KTP',
  "id_card_number" VARCHAR(50) NOT NULL DEFAULT '',
  "id_card_file" VARCHAR(255) NOT NULL DEFAULT '',
  "first_name" VARCHAR(20) NOT NULL,
  "last_name" VARCHAR(20) NOT NULL,
  "phone_number" VARCHAR(13) NOT NULL,
  "email" VARCHAR(100) UNIQUE NOT NULL,
  "username" VARCHAR(100) UNIQUE NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "status" customer_enum NOT NULL DEFAULT 'pending',
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()',
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01'
);

CREATE TABLE "m_account" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "customer_id" UUID UNIQUE NOT NULL,
  "number" VARCHAR(12) UNIQUE NOT NULL,
  "balance" NUMERIC(10, 2) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "m_merchant" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" VARCHAR(200) UNIQUE NOT NULL,
  "balance" NUMERIC(10, 2) NOT NULL DEFAULT 0,
  "address" VARCHAR(255) NOT NULL,
  "website" VARCHAR(255) UNIQUE NOT NULL,
  "email" VARCHAR(100) UNIQUE NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()',
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01'
);

CREATE TABLE "transaction_history" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "transaction_type" transaction_type NOT NULL,
  "from_account_id" UUID NOT NULL,
  "to_account_id" UUID,
  "to_merchant_id" UUID,
  "amount" NUMERIC(10, 2) NOT NULL,
  "description" VARCHAR(255) NOT NULL DEFAULT '',
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "entries" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "account_id" UUID NOT NULL,
  "transaction_type" transaction_type NOT NULL,
  "transaction_id" UUID NOT NULL,
  "amount" NUMERIC(10, 2) NOT NULL
);

CREATE INDEX ON "m_customer" ("username");

CREATE INDEX ON "m_customer" ("id_card_number");

CREATE INDEX ON "m_customer" ("email");

CREATE INDEX ON "m_customer" ("phone_number");

CREATE INDEX ON "m_account" ("customer_id");

CREATE INDEX ON "m_account" ("number");

CREATE INDEX ON "m_merchant" ("name");

CREATE INDEX ON "m_merchant" ("email");

CREATE INDEX ON "m_merchant" ("address");

CREATE INDEX ON "transaction_history" ("from_account_id");

CREATE INDEX ON "transaction_history" ("to_account_id");

CREATE INDEX ON "transaction_history" ("to_merchant_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "m_account" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("transaction_id") REFERENCES "transaction_history" ("id");

ALTER TABLE "m_account" ADD FOREIGN KEY ("customer_id") REFERENCES "m_customer" ("id");

ALTER TABLE "transaction_history" ADD FOREIGN KEY ("from_account_id") REFERENCES "m_account" ("id");

ALTER TABLE "transaction_history" ADD FOREIGN KEY ("to_account_id") REFERENCES "m_account" ("id");

ALTER TABLE "transaction_history" ADD FOREIGN KEY ("to_merchant_id") REFERENCES "m_merchant" ("id");

-- create admin customer/user
    -- "username" : "admin",
    -- "password" : "admin123Dev"
INSERT INTO m_customer
("role","first_name", "last_name", "phone_number", "email", "username", "password", "status")
VALUES
('admin','dev','admin','080088000808','admin@dev.com','admin','$2a$10$26HYIguUqwVFmAEtETpCc.0bcLAjjM6vgJCWnTDvqvmqRjIzTsAye', 'active');