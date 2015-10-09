package viewmodels

import (
)


type Home struct {
	Title string
	Active string
}


func GetHome() Home {

	result := Home{
		Title: "Lemonade stand society",
		Active: "home",
	}

	return result
	
}