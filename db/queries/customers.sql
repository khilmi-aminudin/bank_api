-- name: CreateCustomer :one
INSERT INTO m_customer (
    "id_card_type",
    "id_card_number",
    "id_card_file",
    "first_name",
    "last_name",
    "phone_number",
    "email",
    "username",
    "password"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetAllCustomers :many
SELECT * FROM m_customer
LIMIT $1 OFFSET $2;

-- name: GetCustomerById :one
SELECT * FROM m_customer WHERE "id" = $1;

-- name: GetCustomerByEmail :one
SELECT "id", "role", "username", "email", "password", "status" FROM m_customer WHERE email = $1;

-- name: GetCustomerByUsername :one
SELECT "id", "role", "username", "email", "password", "status" FROM m_customer WHERE username = $1;

-- name: UpdateCustomer :one
UPDATE m_customer
SET "id_card_type" = $2, "id_card_number" = $3, "id_card_file" = $4, "status" = $5
WHERE "id" = $1 RETURNING *;