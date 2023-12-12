CREATE TABLE "users" (
  "email" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "itemsGroupKey" varchar NOT NULL
);

CREATE TABLE "items" (
  "id" bigserial PRIMARY KEY,
  "description" varchar UNIQUE NOT NULL,
  "groupKey" varchar NOT NULL,
  "user" varchar NOT NULL
);

CREATE INDEX ON "items" ("groupKey");

ALTER TABLE "items" ADD FOREIGN KEY ("groupKey") REFERENCES "users" ("itemsGroupKey");

ALTER TABLE "items" ADD FOREIGN KEY ("user") REFERENCES "users" ("email");
DROP TABLE IF EXISTS "items";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "groups";