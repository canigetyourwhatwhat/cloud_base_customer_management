package entity

import "time"

type Customer struct {
	ID          string    `json:"id" db:"id"`
	CustomerID  string    `json:"customerID" db:"customer_id"`
	CompanyName string    `json:"companyName" db:"company_name"`
	FirstName   string    `json:"firstName" db:"first_name"`
	LastName    string    `json:"lastName" db:"last_name"`
	Email       string    `json:"email"`
	UpdatedAt   time.Time `db:"updated_at"`
	CreatedAt   time.Time `db:"created_at"`
}
