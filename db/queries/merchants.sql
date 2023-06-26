-- name: CreateMerchant :one
INSERT INTO m_merchant (
    "name",
    "address",
    "website",
    "email"
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateMerchant :exec
UPDATE m_merchant 
SET "name" = $2, "address" = $3, "website" = $4
WHERE "id" = $1;

-- name: GetAllMerchants :many
SELECT * FROM m_merchant;

-- name: GetMerchantByName :one
SELECT * FROM m_merchant
WHERE "name" = $1;

-- name: GetMerchantById :one
SELECT * FROM m_merchant
WHERE "id" = $1;

-- name: AddMerchantBalance :one
UPDATE m_merchant
SET "balance" = balance + sqlc.arg(balance)
WHERE "id" = sqlc.arg(id)
RETURNING *;
