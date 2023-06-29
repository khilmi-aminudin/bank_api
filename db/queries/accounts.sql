-- name: CreateAccount :one
INSERT INTO m_account(
    "customer_id",
    "number",
    "balance"
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: AddAccountBalance :one
UPDATE m_account
SET balance = balance + sqlc.arg(balance)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: GetAccountByNumber :one
SELECT * FROM m_account 
WHERE "number" = $1;

-- name: GetAccountByCustomerID :one
SELECT * FROM m_account 
WHERE "customer_id" = $1;