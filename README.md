# grpc-go-htmx-graphql

microservice impl using various technologies

## ui-server

The ui that communicates with graphql-server

* Bootstrapped with go-blueprint
* UI components with [MDB](https://mdbootstrap.com/)

`go-blueprint create --name ui-server --framework gorilla/mux --driver none --advanced --feature htmx --git skip`

## graphql-server

The graphql server that serves data to the ui and fetches data from the microservices

* Bootstrapped with [gqlgen](https://gqlgen.com/)

## customer-service

gRPC service with a SQLite database that holds the customer domain

## order-service

gRPC service with a SQLite database that holds the order domain


