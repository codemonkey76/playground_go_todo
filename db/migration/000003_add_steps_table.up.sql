CREATE TABLE IF NOT EXISTS "steps" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "task_id" BIGINT NOT NULL,
  "name" VARCHAR NOT NULL,
  "completed" BOOLEAN NOT NULL DEFAULT false,

  -- Timestamps
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,

  -- Foreign keys
  FOREIGN KEY ("task_id") REFERENCES "tasks" ("id") ON DELETE CASCADE
);

