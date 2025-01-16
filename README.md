# grpc-go-htmx-graphql

microservice impl using various technologies - golang project with workspace

## ui-server

The ui that communicates with graphql-server

``
go work init ./customer-service ./graphql-server ./order-service ./product-service ./shared
``

* Next.js
* MUI

## graphql-server

The graphql server that serves data to the ui and fetches data from the microservices

* Bootstrapped with [gqlgen](https://gqlgen.com/)

## customer-service

Unary gRPC service with a SQLite database that holds the customer domain

## order-service

Unary gRPC service with a SQLite database that holds the order domain

## product-service

Unary gRPC service with a SQLite database that holds the product domain

## shared

library with shared go files for the project


