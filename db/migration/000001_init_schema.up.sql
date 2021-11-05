CREATE TABLE "users" (
    "user_id"       BIGSERIAL PRIMARY KEY,
    "username"      VARCHAR(40) NOT NULL,
    "password"      VARCHAR(40) NOT NULL,
    "email"         VARCHAR(256),
    "created_at"    TIMESTAMP NOT NULL DEFAULT(now())
);