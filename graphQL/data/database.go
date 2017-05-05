package data

import "gopkg.in/mgo.v2/bson"

type Quote struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Text   string        `json:"text"`
	Author string        `json:"author"`
}

type Quote2 struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

type QuoteList []Quote

type QuotesLibraryNode struct {
	allQuotes QuoteList
}

func GetQuoteByID(id string) interface{} {
	return Quote{}
}
