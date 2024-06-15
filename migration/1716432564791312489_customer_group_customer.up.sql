CREATE TABLE IF NOT EXISTS "customer_group_customer" (
"created_by" text NULL,
"id" text NOT NULL PRIMARY KEY,
"customer_id" text NOT NULL,
"customer_group_id" text NOT NULL,
"metadata" jsonb NULL,
"created_at" timestamptz NOT NULL,
"updated_at" timestamptz NOT NULL,
CONSTRAINT "customer_group_customer_customer_group_id_foreign" FOREIGN KEY ("customer_group_id") REFERENCES "customer_group" ("id") ON DELETE cascade,
CONSTRAINT "customer_group_customer_customer_id_foreign" FOREIGN KEY ("customer_id") REFERENCES "customer" ("id") ON DELETE cascade
);
