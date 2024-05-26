package models

type Invoice struct {
	ID              int
	Amount          int    `json:"amount"`
	Department      string `json:"department"`
	ManagerApproval bool   `json:"manager_approval"`
}
