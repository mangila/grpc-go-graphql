run:
	@echo "run graphql-server"
	@go run server.go
build:
	@echo "building graphql-server"
	@go build -o main server.go
generate-schema:
	@echo "generate-schema graphql-server"
	@go run github.com/99designs/gqlgen generate
clean:
	@echo "clean graphql-server"
	@rm -f main