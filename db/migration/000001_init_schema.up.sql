CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "links" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "address" varchar NOT NULL,
  "user_id" bigint NOT NULL
);

ALTER TABLE "links" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");