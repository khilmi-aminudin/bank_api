// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CustomerEnum string

const (
	CustomerEnumActive   CustomerEnum = "active"
	CustomerEnumInactive CustomerEnum = "inactive"
	CustomerEnumPending  CustomerEnum = "pending"
	CustomerEnumBlocked  CustomerEnum = "blocked"
)

func (e *CustomerEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CustomerEnum(s)
	case string:
		*e = CustomerEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for CustomerEnum: %T", src)
	}
	return nil
}

type NullCustomerEnum struct {
	CustomerEnum CustomerEnum
	Valid        bool // Valid is true if CustomerEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCustomerEnum) Scan(value interface{}) error {
	if value == nil {
		ns.CustomerEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CustomerEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCustomerEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CustomerEnum), nil
}

type IDCardType string

const (
	IDCardTypeKTP      IDCardType = "KTP"
	IDCardTypeSIM      IDCardType = "SIM"
	IDCardTypePassport IDCardType = "Passport"
)

func (e *IDCardType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = IDCardType(s)
	case string:
		*e = IDCardType(s)
	default:
		return fmt.Errorf("unsupported scan type for IDCardType: %T", src)
	}
	return nil
}

type NullIDCardType struct {
	IDCardType IDCardType
	Valid      bool // Valid is true if IDCardType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullIDCardType) Scan(value interface{}) error {
	if value == nil {
		ns.IDCardType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.IDCardType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullIDCardType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.IDCardType), nil
}

type TransactionStatus string

const (
	TransactionStatusSuccess TransactionStatus = "success"
	TransactionStatusFailed  TransactionStatus = "failed"
)

func (e *TransactionStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TransactionStatus(s)
	case string:
		*e = TransactionStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for TransactionStatus: %T", src)
	}
	return nil
}

type NullTransactionStatus struct {
	TransactionStatus TransactionStatus
	Valid             bool // Valid is true if TransactionStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTransactionStatus) Scan(value interface{}) error {
	if value == nil {
		ns.TransactionStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TransactionStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTransactionStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TransactionStatus), nil
}

type TransactionType string

const (
	TransactionTypeTopup      TransactionType = "topup"
	TransactionTypeWithdrawal TransactionType = "withdrawal"
	TransactionTypeTransfer   TransactionType = "transfer"
	TransactionTypePayment    TransactionType = "payment"
)

func (e *TransactionType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TransactionType(s)
	case string:
		*e = TransactionType(s)
	default:
		return fmt.Errorf("unsupported scan type for TransactionType: %T", src)
	}
	return nil
}

type NullTransactionType struct {
	TransactionType TransactionType
	Valid           bool // Valid is true if TransactionType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTransactionType) Scan(value interface{}) error {
	if value == nil {
		ns.TransactionType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TransactionType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTransactionType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TransactionType), nil
}

type MAccount struct {
	ID         uuid.UUID   `json:"id"`
	CustomerID uuid.UUID   `json:"customer_id"`
	Number     string      `json:"number"`
	Balance    interface{} `json:"balance"`
	CreatedAt  time.Time   `json:"created_at"`
}

type MCustomer struct {
	ID           uuid.UUID    `json:"id"`
	IDCardType   IDCardType   `json:"id_card_type"`
	IDCardNumber string       `json:"id_card_number"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	PhoneNumber  string       `json:"phone_number"`
	Email        string       `json:"email"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	Status       CustomerEnum `json:"status"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type MMerchant struct {
	ID        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	Balance   interface{} `json:"balance"`
	Address   string      `json:"address"`
	Website   string      `json:"website"`
	Email     string      `json:"email"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type TransactionHistory struct {
	ID              uuid.UUID         `json:"id"`
	TransactionType TransactionType   `json:"transaction_type"`
	FromAccountID   uuid.UUID         `json:"from_account_id"`
	ToAccountID     uuid.NullUUID     `json:"to_account_id"`
	ToMerchantID    uuid.NullUUID     `json:"to_merchant_id"`
	Amount          interface{}       `json:"amount"`
	Description     string            `json:"description"`
	Status          TransactionStatus `json:"status"`
	CreatedAt       time.Time         `json:"created_at"`
}
