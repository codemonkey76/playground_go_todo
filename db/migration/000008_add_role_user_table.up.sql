CREATE TABLE IF NOT EXISTS "role_user" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "role_id" BIGINT NOT NULL,
  "user_id" BIGINT NOT NULL,

  -- Timestamps
  "created_at" TIMESTAMP DEFAULT (now()),
  "updated_at" TIMESTAMP,

  -- Foreign keys
  FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON DELETE CASCADE,
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

