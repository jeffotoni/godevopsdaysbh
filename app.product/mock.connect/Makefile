# Makefile

# postgres in docker
# export DB_HOST := 192.0.0.3
# export DB_USER := gopher
# export DB_PASSWORD := 1234
# export DB_BANCO := gopher
# export DB_PORT := 5432

# postgres localhost
export DB_HOST := 172.17.0.2
export DB_USER := gopher
export DB_PASSWORD := 12345
export DB_BANCO := apidevopsdaysbh
export DB_PORT := 5432

loginfind:
	go run main-create-find-login.go

loginemail:
	go run main-connect-login-email.go

loginconnect:
	go run main-connect-login.go

loginconnect2:
	go run main-connect-login2.go

empresaconnect:
	go run main-connect-empresa.go

empresainsert:
	go run main-connect-insert-empresa.go