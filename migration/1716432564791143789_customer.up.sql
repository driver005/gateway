CREATE TABLE IF NOT EXISTS "customer" (
"metadata" jsonb NULL,
"created_at" timestamptz NOT NULL,
"deleted_at" timestamptz NULL,
"created_by" text NULL,
"email" text NULL,
"has_account" boolean NOT NULL,
"first_name" text NULL,
"last_name" text NULL,
"phone" text NULL,
"updated_at" timestamptz NOT NULL,
"id" text NOT NULL PRIMARY KEY,
"company_name" text NULL
);
