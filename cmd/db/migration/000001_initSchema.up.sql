CREATE TABLE "User" (
  "id" integer PRIMARY KEY,
  "username" varchar,
  "created_at" timestamp
);

CREATE TABLE "Files" (
  "id" integer PRIMARY KEY,
  "title" varchar,
  "body" bytea,
  "status" varchar,
  "created_at" timestamp
);

ALTER TABLE "User" ADD FOREIGN KEY ("id") REFERENCES "Files" ("id");