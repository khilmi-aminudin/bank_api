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