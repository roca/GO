package data

type Quote struct {
	//ID     int    `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

type QuoteList []Quote
