# VARIABLES	

MODULE = $(shell go list -m)
PACKAGES := $(shell go list ./... | grep -v /vendor/)


.PHONY: test
test: ## run unit tests
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: clean
clean: ## remove temporary files
	rm -rf server coverage.out coverage-all.out


.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	go tool cover -html=coverage-all.out


.PHONY: run
run: ## run the API server
	go run cmd/main.go


.PHONY: lint
lint: ## run linter
	golangci-lint run

.PHONY: compose
compose: ## start the database
	docker-compose -f deployments/docker-compose.yml up -d


