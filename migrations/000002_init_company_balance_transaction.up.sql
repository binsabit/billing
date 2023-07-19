CREATE TABLE IF NOT EXISTS  company_balance_transactions(
    id bigint not null auto_increment primary key,
    company_id bigint not null,
    type int not null,
    amount numeric(12,2) not null default 0,
    data json,
    sender_type smallint not null,
    sender_id bigint,
    provider_id smallint not null,
    status smallint not null,
    is_init bool not null default false,
    paid_at timestamp
);