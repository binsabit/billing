-- name: CreateTransaction :execresult
insert into company_balance_transactions(
    company_id,    type,    sender_type,
    sender_id,    unsigned(amount),    status,    provider_id
) values (
    ?,?,?,?,?,?,?
);

-- name: UpdateTransaction :execresult


