CREATE TABLE IF NOT EXISTS company_cart_services(
    id bigint not null auto_increment  primary key,
    company_cart_id bigint not null,
    service_id bigint not null,
    foreign key (company_cart_id) references company_carts(id),
    foreign key (service_id) references services(id)
);