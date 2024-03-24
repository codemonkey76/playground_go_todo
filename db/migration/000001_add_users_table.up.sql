CREATE TABLE IF NOT EXISTS "users" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" VARCHAR NOT NULL,
  "email" VARCHAR NOT NULL UNIQUE,
  "email_verified_at" TIMESTAMP,
  "password" VARCHAR NOT NULL,
  "remember_token" VARCHAR,

    -- Timestamps
  "created_at" TIMESTAMP DEFAULT (now()),
  "updated_at" TIMESTAMP
);

