CREATE TYPE "roles" AS ENUM (
  'guard',
  'forward',
  'center'
);

CREATE TABLE "players" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "role" roles NOT NULL,
  "team" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "coaches" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "teams" (
  "id" bigserial PRIMARY KEY,
  "coach_id" bigint NOT NULL,
  "wins" integer NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "player" bigint NOT NULL,
  "from_team" bigint NOT NULL,
  "to_team" bigint NOT NULL
);

CREATE INDEX ON "players" ("name");

CREATE INDEX "roles_in_team" ON "players" ("role", "team");

CREATE INDEX ON "coaches" ("username");

CREATE INDEX ON "teams" ("coach_id");

ALTER TABLE "players" ADD FOREIGN KEY ("team") REFERENCES "teams" ("id");

ALTER TABLE "teams" ADD FOREIGN KEY ("coach_id") REFERENCES "coaches" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("player") REFERENCES "players" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_team") REFERENCES "teams" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_team") REFERENCES "teams" ("id");
