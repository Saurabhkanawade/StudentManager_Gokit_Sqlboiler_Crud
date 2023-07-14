NAME := customer-patient-service

VERSION := v1.0.0

LDFLAGS  := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

NOVENDOR := $(shell go list ./... | grep -v vendor)

TEST_PATTERN?=.

TEST_OPTIONS?=-race -covermode=atomic -coverprofile=coverage.txt

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME): $(GO_SRC_DIRS)
	go build $(LDFLAGS) -buildvcs=false -o bin/$(NAME)

.PHONY: setup
setup: init-git-hooks  ## Installs all dependencies
	go mod tidy -compat=1.17
	which sqlboiler || (go install github.com/volatiletech/sqlboiler/v4@latest && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest)

.PHONY: install
install:  ## runs install
	go install $(LDFLAGS)

.PHONY: test
test: ## runs test with coverage (doesn't produce a report)
	go test -v -coverpkg=./... -coverprofile=bin/coverage.txt ./...

.PHONY: test-no-cover
test-no-cover: ## runs test without coverage
	go test -v ./...

.PHONY: cover
cover: test ## Run all the tests and opens the coverage report
		go tool cover -html=bin/coverage.txt

.PHONY: fmt
fmt: ## runs fmt
	go fmt $(NOVENDOR)
.PHONY: clean
clean: ## clean up
	rm -rf bin/*

.PHONY: code-gen-db
code-gen-db:  ## generates code from db/sqlboiler.toml
	sqlboiler psql
	grep -rl "DeletedAt" internal/dbmodels | xargs sed -i "" 's/DeletedAt/DeletedTS/g'  # fixes sqlboiler bug
	grep -rl "deleted_at" internal/dbmodels | xargs sed -i "" 's/deleted_at/deleted_ts/g'

.PHONY: help
help:  ## displays this message
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:  ## displays the version
	@echo $(VERSION)

install-mockery: ## installs mockery through brew.
	which mockery || brew install mockery

update-mockery: install-mockery ## updates mockery through brew.
	brew upgrade mockery

mocks: ## regenerates mockery mocks
	mockery --all --outpkg=mocks --output=./mocks --inpackage --dir=./internal/

install-lint: ## installs golangci-lint linter
	which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1

lint: ## runs golangci-lint linter
	golangci-lint run --config golangci.yml --timeout 2m0s

swagger-install: ## install go swagger tool
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.3

swagger-generate: ## generate swagger.json file
	swagger generate spec -o ./swagger-ui/swagger.json --scan-models

swagger-validate: swagger-generate ## validate the swagger docs
	swagger validate ./swagger-ui/swagger.json

swagger-serve: swagger-generate ## host local swagger page for debugging
		swagger serve ./swagger-ui/swagger.json -F swagger

run: ## run the web server locally
	 go build -o bin/customer-patient-service && ./bin/customer-patient-service

.PHONY: pre-commit
pre-commit: test-no-cover swagger-validate lint   ## call before committing and pushing changes
	@echo 'pre-commit checks completed without error. You are clear to commit to your branch'

init-git-hooks: ## sets the git hooks folder for
	git config core.hooksPath ./git-hooks

.PHONY: vendor
vendor: ## run go mod tidy and vendor
	go mod tidy -compat=1.19 && go mod vendor
