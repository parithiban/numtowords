.DEFAULT_GOAL := explain
PROJECTNAME := numtowords
.PHONY: explain
explain:
	### Welcome
	#
	#   
	#     _      _     _      _____  ____  _      ____  ____  ____  ____ 
	#    / \  /|/ \ /\/ \__/|/__ __\/  _ \/ \  /|/  _ \/  __\/  _ \/ ___\ 
	#    | |\ ||| | ||| |\/||  / \  | / \|| |  ||| / \||  \/|| | \||    \ 
	#    | | \||| \_/|| |  ||  | |  | \_/|| |/\||| \_/||    /| |_/|\___ |
	#    \_/  \|\____/\_/  \|  \_/  \____/\_/  \|\____/\_/\_\\____/\____/
	#                                                                                                      
	#
	### Installation
	#
	### Targets
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: ## install all project dependency
	go get ./...

.PHONY: vet
vet: ## Vet the code
	go vet ./...

.PHONY: lint
lint: ## Lint the code
	golint -set_exit_status ./...

.PHONY: security-check
security-check: ## Inspect code for security vulnerabilities
	gosec ./...

.PHONY: cyclo
cyclo-check: ## Inspect code for cyclomatic complexities
	gocyclo -over 15 .

.PHONY: build
build: ## Build the helper
	go build -o numtowords numtowords.go

.PHONY: test
test: ## Run the unit tests
	go test ./... -coverprofile=coverage.out

.PHONY: test-coverage
test-cov: test ## Run the unit tests with coverage
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
