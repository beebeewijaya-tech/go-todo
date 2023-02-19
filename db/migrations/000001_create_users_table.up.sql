CREATE TABLE users (
    "id" bigserial PRIMARY KEY,
    "email" varchar UNIQUE NOT NULL,
    "fullname" varchar NOT NULL,
    "password" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (NOW())
);