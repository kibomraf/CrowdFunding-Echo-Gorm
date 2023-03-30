-- Active: 1679307292556@@127.0.0.1@5432@crowdfunding
CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "occupation" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password_hash" varchar NOT NULL,
  "avatar_file_name" varchar,
  "role" varchar NOT NULL,
  "token_vuejs" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "campaign" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint,
  "name" varchar NOT NULL,
  "backer_account" bigint DEFAULT 0,
  "short_description" varchar NOT NULL,
  "goal_ammount" bigint NOT NULL,
  "current_ammount" bigint DEFAULT 0,
  "perks" text,
  "slug" varchar,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "campaignimage" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "campaign_id" bigint,
  "file_name" varchar NOT NULL,
  "isPrimary" boolean NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "transaction" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "campaign_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "ammount" bigint NOT NULL,
  "status" varchar NOT NULL,
  "code_transaction" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

ALTER TABLE "campaign" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "campaignimage" ADD FOREIGN KEY ("campaign_id") REFERENCES "campaign" ("id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("campaign_id") REFERENCES "campaign" ("id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");