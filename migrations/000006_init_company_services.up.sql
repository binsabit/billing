CREATE TABLE IF NOT EXISTS company_services(
    id bigint not null auto_increment primary key,
    company_id bigint not null,
    service_id bigint not null,
    hash_code varchar(255),
    employee_id bigint not null,
    amount numeric(12, 2) not null default 0,
    discount numeric(12, 2) not null default 0,
    item_limit int,
    employee_limit int,
    stock_limit int,
    extra_stock_id bigint,
    activated_at timestamp null default null,
    expires_at timestamp null default null,
    created_at timestamp null default null,
    foreign key (service_id) references services(id)
);