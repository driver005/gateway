CREATE TABLE IF NOT EXISTS "customer_group" (
"created_by" text NULL,
"created_at" timestamptz NOT NULL,
"updated_at" timestamptz NOT NULL,
"deleted_at" timestamptz NULL,
"id" text NOT NULL PRIMARY KEY,
"name" text NULL,
"metadata" jsonb NULL
);
