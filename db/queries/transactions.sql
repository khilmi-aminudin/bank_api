-- name: CreateTransactionHistory :one
INSERT INTO transaction_history (
    "transaction_type",
    "from_account_id",
    "to_account_id",
    "to_merchant_id",
    "amount",
    "description"
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetTransactionHistory :many
SELECT * FROM transaction_history 
WHERE from_account_id = $1;

-- name: GetTransactionHistoryByType :many
SELECT * FROM transaction_history 
WHERE transaction_type = $1 AND from_account_id = $2;

