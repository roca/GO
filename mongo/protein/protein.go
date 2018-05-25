package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Protein struct {
	Id       bson.ObjectId `bson:"_id"`
	RegnName string        `bson:"regnName"`
}

func main() {
	session, err := mgo.Dial("tarbioinformatics01")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("Bioregistry").C("protein")

	results := []Protein{}
	err = c.Find(bson.M{"_type": "protein"}).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(results.regnName)

	for i, p := range results {
		fmt.Println(i, p)
	}

}
