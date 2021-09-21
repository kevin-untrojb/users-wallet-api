test-all: dependencies format imports mocking testing
clean : clean-docker remove-db
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
testing:
	@echo "Running tests"
	@go test ./... -covermode=atomic -coverpkg=./... -count=1 -race
build-api:
	@echo "Starting docker containers and application set up"
	@docker-compose -f build/docker-compose.yml up --build api mysqldb
clean-docker:
	@echo "Cleaning docker containers"
	@docker-compose -f build/docker-compose.yml down -v
remove-db:
	@echo "Removing db folder"
	@rm -r build/mysql/db
