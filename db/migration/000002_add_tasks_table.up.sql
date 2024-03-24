CREATE TABLE IF NOT EXISTS "tasks" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "owner_id" BIGINT NOT NULL,
  "assignee_id" BIGINT,
  "name" VARCHAR NOT NULL,
  "completed" BOOLEAN NOT NULL DEFAULT false,

    -- Timestamps
  "created_at" TIMESTAMP DEFAULT (now()),
  "updated_at" TIMESTAMP,

    -- Forreign key
    FOREIGN KEY("owner_id") REFERENCES "users"("id"),
    FOREIGN KEY("assignee_id") REFERENCES "users"("id") ON DELETE SET NULL
);

