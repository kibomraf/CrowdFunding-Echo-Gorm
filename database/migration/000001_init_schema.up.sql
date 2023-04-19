-- Active: 1680084304820@@127.0.0.1@5432@crowdfunding
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

CREATE TABLE "campaigns" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint,
  "name" varchar NOT NULL,
  "backer_account" bigint DEFAULT 0,
  "short_description" varchar NOT NULL,
  "description" varchar NOT NULL,
  "goal_amount" bigint NOT NULL,
  "current_amount" bigint DEFAULT 0,
  "perks" text,
  "slug" varchar,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "campaign_images" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "campaign_id" bigint,
  "file_name" varchar NOT NULL,
  "is_primary" boolean NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "campaign_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "status" varchar NOT NULL,
  "code_transaction" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

ALTER TABLE "campaigns" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "campaign_images" ADD FOREIGN KEY ("campaign_id") REFERENCES "campaigns" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("campaign_id") REFERENCES "campaigns" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");