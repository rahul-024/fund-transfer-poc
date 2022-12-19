CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "currency" varchar NOT NULL,
  "owner" varchar NOT NULL
);