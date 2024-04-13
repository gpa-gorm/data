help:
	@echo 'Gpa-Gorm CLI Makefile.'
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo 'Targets:'
	@echo 'test: run tests'
	@echo 'format: format code'
	@echo 'tidy: configures the go.mod file'
	@echo 'verify: verify dependencies'
	@echo 'vendor: download dependencies for offline use'
	@echo 'serve-pkg: serve the package documentation'
	@echo 'serve-docs: serve the documentation'

test:
	go clean -cache && go test -v ./pkg/tests

format:
	gofmt -w -s .

tidy:
	go mod tidy

verify:
	go mod verify
	
vendor:
	go mod vendor

serve-pkg:
	pkgsite -open .

serve-docs:
	godoc -http=:6060