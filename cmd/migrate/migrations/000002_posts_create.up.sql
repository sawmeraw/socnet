CREATE TABLE IF NOT EXISTS posts(
    id bigserial primary key,
    title text not null,
    user_id bigserial not null,
    content text not null,
    created_at timestamp(0) with time zone not null default now()
);