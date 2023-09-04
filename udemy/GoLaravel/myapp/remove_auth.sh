 #!/usr/bin/env bash
 
./celeritas migrate down
./celeritas migrate down
rm data/token.go
rm data/user.go
rm data/remember_token.go
rm handlers/auth-handlers.go
rm mail/password-reset.html.tmpl
rm mail/password-reset.plain.tmpl
rm middleware/auth-token.go
rm middleware/auth.go
rm middleware/remember.go
rm migrations/*_auth.postgres.down.sql
rm migrations/*_auth.postgres.up.sql
rm views/forgot.jet
rm views/login.jet
rm views/reset-password.jet
