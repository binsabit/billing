CREATE TABLE IF NOT EXISTS company_balances(
    id bigint not null  auto_increment,
    company_id bigint not null,
    balance numeric(12, 2) not null default 0,
    primary key (id)
);