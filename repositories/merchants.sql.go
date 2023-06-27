// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: merchants.sql

package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addMerchantBalance = `-- name: AddMerchantBalance :one
UPDATE m_merchant
SET "balance" = balance + $1
WHERE "id" = $2
RETURNING id, name, balance, address, website, email, created_at, updated_at
`

type AddMerchantBalanceParams struct {
	Balance float64   `json:"balance"`
	ID      uuid.UUID `json:"id"`
}

func (q *Queries) AddMerchantBalance(ctx context.Context, arg AddMerchantBalanceParams) (MMerchant, error) {
	row := q.db.QueryRowContext(ctx, addMerchantBalance, arg.Balance, arg.ID)
	var i MMerchant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.Address,
		&i.Website,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createMerchant = `-- name: CreateMerchant :one
INSERT INTO m_merchant (
    "name",
    "address",
    "website",
    "email"
) VALUES (
    $1, $2, $3, $4
) RETURNING id, name, balance, address, website, email, created_at, updated_at
`

type CreateMerchantParams struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Website string `json:"website"`
	Email   string `json:"email"`
}

func (q *Queries) CreateMerchant(ctx context.Context, arg CreateMerchantParams) (MMerchant, error) {
	row := q.db.QueryRowContext(ctx, createMerchant,
		arg.Name,
		arg.Address,
		arg.Website,
		arg.Email,
	)
	var i MMerchant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.Address,
		&i.Website,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllMerchants = `-- name: GetAllMerchants :many
SELECT id, name, balance, address, website, email, created_at, updated_at FROM m_merchant
`

func (q *Queries) GetAllMerchants(ctx context.Context) ([]MMerchant, error) {
	rows, err := q.db.QueryContext(ctx, getAllMerchants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MMerchant{}
	for rows.Next() {
		var i MMerchant
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Balance,
			&i.Address,
			&i.Website,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMerchantById = `-- name: GetMerchantById :one
SELECT id, name, balance, address, website, email, created_at, updated_at FROM m_merchant
WHERE "id" = $1
`

func (q *Queries) GetMerchantById(ctx context.Context, id uuid.UUID) (MMerchant, error) {
	row := q.db.QueryRowContext(ctx, getMerchantById, id)
	var i MMerchant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.Address,
		&i.Website,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getMerchantByName = `-- name: GetMerchantByName :one
SELECT id, name, balance, address, website, email, created_at, updated_at FROM m_merchant
WHERE "name" = $1
`

func (q *Queries) GetMerchantByName(ctx context.Context, name string) (MMerchant, error) {
	row := q.db.QueryRowContext(ctx, getMerchantByName, name)
	var i MMerchant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.Address,
		&i.Website,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateMerchant = `-- name: UpdateMerchant :exec
UPDATE m_merchant 
SET "name" = $2, "address" = $3, "website" = $4, "updated_at" = $5
WHERE "id" = $1
`

type UpdateMerchantParams struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Website   string    `json:"website"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) UpdateMerchant(ctx context.Context, arg UpdateMerchantParams) error {
	_, err := q.db.ExecContext(ctx, updateMerchant,
		arg.ID,
		arg.Name,
		arg.Address,
		arg.Website,
		arg.UpdatedAt,
	)
	return err
}