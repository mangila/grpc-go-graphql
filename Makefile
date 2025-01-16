# Root Makefile

ORDER_SERVICE_DIR := order-service
CUSTOMER_SERVICE_DIR := customer-service
PRODUCT_SERVICE_DIR := product-service
GRAPHQL_SERVER_DIR := graphql-server

# Apply k8s definitions
minikube-apply-resources:
	minikube kubectl -- apply -f scripts/config-k8s.yml
	minikube kubectl -- apply -f customer-service/k8s.yml
	minikube kubectl -- apply -f graphql-server/k8s.yml
	minikube kubectl -- apply -f order-service/k8s.yml
	minikube kubectl -- apply -f product-service/k8s.yml

# run from root file
run-order-service:
	@$(MAKE) run -C $(ORDER_SERVICE_DIR)
run-customer-service:
	@$(MAKE) run -C $(CUSTOMER_SERVICE_DIR)
run-product-service:
	@$(MAKE) run -C $(PRODUCT_SERVICE_DIR)
run-graphql-server:
	 @$(MAKE) run -C $(GRAPHQL_SERVER_DIR)

# build from root file
build-all: build-order-service build-customer-service build-product-service build-graphql-server build-ui-server
build-order-service:
	@$(MAKE) build -C $(ORDER_SERVICE_DIR)
build-customer-service:
	@$(MAKE) build -C $(CUSTOMER_SERVICE_DIR)
build-product-service:
	@$(MAKE) build -C $(PRODUCT_SERVICE_DIR)
build-graphql-server:
	@$(MAKE) build -C $(GRAPHQL_SERVER_DIR)

# clean from root file
clean-all: clean-order-service clean-customer-service clean-product-service clean-graphql-server clean-ui-server
clean-order-service:
	@$(MAKE) clean -C $(ORDER_SERVICE_DIR)
clean-customer-service:
	@$(MAKE) clean -C $(CUSTOMER_SERVICE_DIR)
clean-product-service:
	@$(MAKE) clean -C $(PRODUCT_SERVICE_DIR)
clean-graphql-server:
	@$(MAKE) clean -C $(GRAPHQL_SERVER_DIR)

# compile protos
proto-compile-all: proto-compile-service proto-compile-model
proto-compile-service:
	 protoc --proto_path=shared/proto --go_out=. --go-grpc_out=. shared/proto/*/*_service.proto
proto-compile-model:
	protoc --proto_path=shared/proto --go_out=. shared/proto/*/*_model.proto