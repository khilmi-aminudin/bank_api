CREATE TYPE "id_card_type" AS ENUM (
  'KTP',
  'SIM',
  'Passport'
);

CREATE TYPE "transaction_status" AS ENUM (
  'success',
  'failed'
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
  "id" UUID PRIMARY KEY NOT NULL DEFAULT 'uuid_generate_v4()',
  "id_card_type" id_card_type NOT NULL,
  "id_card_number" VARCHAR(50) NOT NULL DEFAULT '',
  "first_name" VARCHAR(20) NOT NULL,
  "last_name" VARCHAR(20) NOT NULL,
  "phone_number" VARCHAR(13) NOT NULL,
  "email" VARCHAR(100) NOT NULL,
  "status" customer_enum NOT NULL DEFAULT 'pending',
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()',
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01'
);

CREATE TABLE "m_account" (
  "id" UUID PRIMARY KEY NOT NULL DEFAULT 'uuid_generate_v4()',
  "customer_id" UUID UNIQUE NOT NULL,
  "number" VARCHAR(12) NOT NULL,
  "balance" "NUMERIC(10, 2)" NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "m_merchant" (
  "id" UUID PRIMARY KEY NOT NULL DEFAULT 'uuid_generate_v4()',
  "name" VARCHAR(200) NOT NULL,
  "balance" "NUMERIC(10, 2)" NOT NULL DEFAULT 0,
  "address" VARCHAR(255) NOT NULL,
  "website" VARCHAR(255) NOT NULL,
  "email" VARCHAR(100) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()',
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01'
);

CREATE TABLE "transaction_history" (
  "id" UUID PRIMARY KEY NOT NULL DEFAULT 'uuid_generate_v4()',
  "transaction_type" transaction_type NOT NULL,
  "from_account_id" UUID NOT NULL,
  "to_account_id" UUID,
  "to_merchant_id" UUID,
  "amount" "NUMERIC(10, 2)" NOT NULL,
  "description" VARCHAR(255) NOT NULL DEFAULT '',
  "status" transaction_status NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

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

ALTER TABLE "m_account" ADD FOREIGN KEY ("customer_id") REFERENCES "m_customer" ("id");

ALTER TABLE "transaction_history" ADD FOREIGN KEY ("from_account_id") REFERENCES "m_account" ("id");

ALTER TABLE "transaction_history" ADD FOREIGN KEY ("to_account_id") REFERENCES "m_account" ("id");

ALTER TABLE "transaction_history" ADD FOREIGN KEY ("to_merchant_id") REFERENCES "m_merchant" ("id");
