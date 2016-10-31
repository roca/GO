package design

import . "github.com/goadesign/goa/design"
import . "github.com/goadesign/goa/design/apidsl"

var _ = API("cellar", func() {
	Description("The wine review service")
	Host("localhost:8080")
})

var BottlePayLoad = Type("BottlePayLoad", func() {
	Description("BottlePayLoad is the type used to create bottles")

	Attribute("name", String, "Name of bottle", func() {
		MinLength(2)
	})

	Attribute("vintage", Integer, "Vintage of bottle", func() {
		Minimum(1900)
	})

	Attribute("rating", Integer, "Name of bottle", func() {
		Minimum(1)
		Maximum(5)
	})

	Required("name", "vintage", "rating")
})
