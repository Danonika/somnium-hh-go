-- +goose Up
-- +goose StatementBegin

CREATE TABLE "users" (
  ID uuid primary key default gen_random_uuid(),
  Created_at timestamptz default now(),
  Email text not null unique,
  hashed_password text not null,
  FIO VARCHAR(255),
  Contact VARCHAR(255),
  resume_title VARCHAR(255),
  resume_link text,
  skills text[]
);

create table "user_roles" (
  "id" int primary key,
  "value" text
);

insert into "user_roles" 
  ("id", "value")
values
  (1, 'user'),
  (2, 'admin')
;

create index users_id_idx on "users" using hash ("id");
create index users_email_idx on "users" using hash ("email");

create table "user_role_bindings" (
  "created_at" timestamptz default now(),
  "user_id" uuid,
  "role_id" int,
  foreign key ("user_id") references "users" ("id") on delete cascade,
  foreign key ("role_id") references "user_roles" ("id") on delete cascade,
  primary key ("user_id", "role_id")
);

create table skills (
  skill VARCHAR(255) not null unique
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "user_role_bindings";

DROP INDEX IF EXISTS users_id_idx;
DROP INDEX IF EXISTS users_email_idx;

DROP TABLE IF EXISTS "user_roles";

DROP TABLE IF EXISTS "users";

DROP TABLE IF EXISTS "skills";
-- +goose StatementEnd
