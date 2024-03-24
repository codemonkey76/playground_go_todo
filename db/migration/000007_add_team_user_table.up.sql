CREATE TABLE IF NOT EXISTS "team_user" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "team_id" BIGINT NOT NULL,
  "user_id" BIGINT NOT NULL,

  -- Timestamps
  "created_at" TIMESTAMP DEFAULT (now()),
  "updated_at" TIMESTAMP,

  -- Foreign Keys
  FOREIGN KEY ("team_id") REFERENCES "teams" ("id"),
  FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

