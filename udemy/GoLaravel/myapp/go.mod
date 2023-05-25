module myapp

go 1.20

replace github.com/roca/celeritas => ../celeritas

require github.com/roca/celeritas v0.0.0-20230522090128-04fd9095989e

require (
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet/v6 v6.2.0 // indirect
	github.com/go-chi/chi/v5 v5.0.8 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)
