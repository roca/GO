package models

type Transfer struct {
	FromAccount string `json:"from" valid:"Required"`
	ToAccount   string `json:"to" valid:"Required"`
	Amount      int    `json:"amount" valid:"Required;Min(0)"`
}
