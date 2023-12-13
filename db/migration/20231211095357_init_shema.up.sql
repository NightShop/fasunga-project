CREATE TABLE "users" (
  "email" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "group_key" varchar NOT NULL
);

CREATE TABLE "items" (
  "id" bigserial PRIMARY KEY,
  "user_email" varchar NOT NULL,
  "description" varchar NOT NULL,
  "group_key" varchar NOT NULL,
  "checked" bool NOT NULL DEFAULT false
);

CREATE TABLE "groups" (
  "group_key" varchar PRIMARY KEY
);

CREATE INDEX ON "items" ("group_key");

ALTER TABLE "users" ADD FOREIGN KEY ("group_key") REFERENCES "groups" ("group_key");

ALTER TABLE "items" ADD FOREIGN KEY ("group_key") REFERENCES "groups" ("group_key");
