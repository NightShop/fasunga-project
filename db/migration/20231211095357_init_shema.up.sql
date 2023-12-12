CREATE TABLE "users" (
  "email" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "group_name" varchar NOT NULL
);

CREATE TABLE "items" (
  "user_email" varchar NOT NULL,
  "item" varchar NOT NULL,
  "group_name" varchar NOT NULL,
  PRIMARY KEY ("group_name", "item")
);

CREATE TABLE "groups" (
  "group_name" varchar PRIMARY KEY
);

CREATE INDEX ON "items" ("group_name");

ALTER TABLE "users" ADD FOREIGN KEY ("group_name") REFERENCES "groups" ("group_name");

ALTER TABLE "items" ADD FOREIGN KEY ("group_name") REFERENCES "groups" ("group_name");
