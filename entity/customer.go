package entity

type Customer struct {
	ID          string `json:"id"`
	CustomerID  string `json:"customerID"`
	CompanyName string `json:"companyName"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
}
