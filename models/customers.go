package models

import m "github.com/khilmi-aminudin/bank_api/repositories"

func NewCustomers(c m.MCustomer) interface{} {
	return struct {
		Status       m.CustomerEnum `json:"status"`
		IdCardType   m.IDCardType   `json:"id_card_type"`
		IdcardNumber string         `json:"id_card_number"`
		IdCardFile   string         `json:"id_card_file"`
		FirstName    string         `json:"first_name"`
		LastName     string         `json:"last_name"`
		PhoneNumber  string         `json:"phone_number"`
		Email        string         `json:"email"`
		Username     string         `json:"username"`
	}{
		Status:       c.Status,
		IdCardType:   c.IDCardType,
		IdcardNumber: c.IDCardNumber,
		IdCardFile:   c.IDCardFile,
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		PhoneNumber:  c.PhoneNumber,
		Email:        c.Email,
		Username:     c.Username,
	}
}
