package middleware

import (
	"myapp/data"

	"github.com/roca/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
