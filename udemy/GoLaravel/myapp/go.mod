module myapp

go 1.20

replace github.com/roca/celeritas => ../celeritas

require github.com/roca/celeritas v0.0.0-20230522090128-04fd9095989e

require (
	github.com/go-chi/chi/v5 v5.0.8 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)
