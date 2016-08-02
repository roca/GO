package models

type Transfer struct {
	FromAccount string `json:"from"`
	ToAccount   string `json:"to"`
	Amount      int    `json:"amount"`
}
