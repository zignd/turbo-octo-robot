package model

// Account models a user account in our bank system
type Account struct {
	ID           string
	FirstName    string `json:"first_name"`
	MiddleName   string `json:"middle_name,required"`
	LastName     string `json:"last_name"`
	BalanceCents uint64 `json:"balance_cents"`
}
