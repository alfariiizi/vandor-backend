-- Create "admin_audit_logs" table
CREATE TABLE "admin_audit_logs" ("id" uuid NOT NULL, "user_email" character varying NOT NULL, "operation" character varying NOT NULL, "model" character varying NOT NULL, "args" jsonb NOT NULL, "result" jsonb NULL, "query" text NULL, "params" text NULL, "source" character varying NOT NULL DEFAULT 'admin-panel', "duration_ms" bigint NULL, "created_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "email" character varying NOT NULL, "first_name" character varying NOT NULL, "last_name" character varying NOT NULL, "password_hash" character varying NOT NULL, "role" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create "products" table
CREATE TABLE "products" ("id" uuid NOT NULL, "name" character varying NOT NULL, "brand" character varying NOT NULL, "category" character varying NOT NULL, "price" double precision NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "creator_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "products_users_products" FOREIGN KEY ("creator_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
