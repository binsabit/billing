CREATE TABLE IF NOT EXISTS services(
    id bigint not null  auto_increment primary key,
    title varchar(100) not null,
    price numeric(8,2) not null,
    type smallint not null,
    month int,
    is_tis bool default false not null,
    is_visible bool default false not null,
    item_limit int,
    employee_limit int,
    stock_limit int,
    deleted_at timestamp
);