fast: swagger-fast
	go run main.go

run: swagger
	go run main.go

swagger:
	go get -u github.com/swaggo/swag/cmd/swag
	swag init

swagger-fast:
	swag init