-- Create "users" table
CREATE TABLE "users" ("id" character varying NOT NULL, "email" character varying NOT NULL, "username" character varying NOT NULL, "first_name" character varying NOT NULL, "last_name" character varying NOT NULL, "phone" character varying NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "user_created_at" to table: "users"
CREATE INDEX "user_created_at" ON "users" ("created_at");
-- Create index "user_email" to table: "users"
CREATE UNIQUE INDEX "user_email" ON "users" ("email");
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create index "users_username_key" to table: "users"
CREATE UNIQUE INDEX "users_username_key" ON "users" ("username");
