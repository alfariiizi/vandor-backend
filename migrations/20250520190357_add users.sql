-- Create "users" table
CREATE TABLE "public"."users" ("id" uuid NOT NULL, "username" text NOT NULL, "password" text NOT NULL, "email" text NOT NULL, "created_at" timestamp NOT NULL DEFAULT now(), "updated_at" timestamp NOT NULL DEFAULT now(), "deleted_at" timestamp NULL, PRIMARY KEY ("id"), CONSTRAINT "users_email_key" UNIQUE ("email"));
