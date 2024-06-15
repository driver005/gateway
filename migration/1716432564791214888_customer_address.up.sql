CREATE TABLE IF NOT EXISTS "customer_address" (
"address_1" text NULL,
"address_2" text NULL,
"province" text NULL,
"postal_code" text NULL,
"is_default_shipping" boolean NOT NULL,
"is_default_billing" boolean NOT NULL,
"last_name" text NULL,
"created_at" timestamptz NOT NULL,
"updated_at" timestamptz NOT NULL,
"address_name" text NULL,
"company" text NULL,
"metadata" jsonb NULL,
"country_code" text NULL,
"phone" text NULL,
"id" text NOT NULL PRIMARY KEY,
"customer_id" text NOT NULL,
"first_name" text NULL,
"city" text NULL,
CONSTRAINT "customer_address_customer_id_foreign" FOREIGN KEY ("customer_id") REFERENCES "customer" ("id") ON UPDATE cascade ON DELETE cascade
);
create unique index "IDX_customer_address_unqiue_customer_billing" on "customer_address" ("customer_id") where "is_default_billing" = true;
create unique index "IDX_customer_address_unique_customer_shipping" on "customer_address" ("customer_id") where "is_default_shipping" = true;
