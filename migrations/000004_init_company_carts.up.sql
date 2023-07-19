CREATE TABLE IF NOT EXISTS company_carts(
    id bigint not null auto_increment primary key,
    company_id bigint not null,
    sender_type smallint not null,
    sender_id bigint,
    provider_id smallint not null,
    status smallint not null,
    sum numeric(12,2) not null default 0,
    paid_at timestamp
);