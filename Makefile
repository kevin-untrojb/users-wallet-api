all: dependencies format imports mocking test
dependencies:
	@echo "Syncing dependencies with go mod tidy"
	@go mod tidy
format:
	@echo "Formatting Go code recursively"
	@go fmt ./...
imports:
	@echo "Executing goimports recursively"
	@goimports -w $(find . -type f -name '*.go') ../
mocking:
	@echo "generating mock files recursively"
	@go generate ./...
test:
	@echo "Running tests"
	@go test ./... -covermode=atomic -coverpkg=./... -count=1 -race
run:
	@echo "Running Application"
	@docker-compose -f build/docker-compose.yml up --build api mysqldb
