# Makefile
# postgres localhost
export DB_HOST := 172.17.0.2
export DB_USER := gopher
export DB_PASSWORD := 12345
export DB_BANCO := apidevopsdaysbh
export DB_PORT := 5432

update:
	rm -f go.*
	go mod init github.com/jeffotoni/godevopsdasybh/app.product
	go run main.go

build:
	go build -ldflags="-s -w" -o app.product main.go
