CREATE TABLE IF NOT EXISTS "permission_role" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "permission_id" BIGINT NOT NULL,
  "role_id" BIGINT NOT NULL,

  -- Timestamps
  "created_at" TIMESTAMP DEFAULT (now()),
  "updated_at" TIMESTAMP,

  -- Foreign keys
  FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id") ON DELETE CASCADE,
  FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON DELETE CASCADE
);

