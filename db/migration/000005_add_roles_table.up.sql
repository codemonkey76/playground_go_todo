CREATE TABLE IF NOT EXISTS "roles" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" VARCHAR NOT NULL,

  -- Timestamps
  "created_at" TIMESTAMP DEFAULT (now()),
  "updated_at" TIMESTAMP
);
