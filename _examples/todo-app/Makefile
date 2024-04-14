help:
	@echo 'Gpa-Gorm CLI Makefile.'
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo 'Targets:'
	@echo 'format: format code'
	@echo 'tidy: configures the go.mod file'
	@echo 'verify: verify dependencies'
	@echo 'vendor: download dependencies for offline use'
	@echo 'serve-pkg: serve the package documentation'
	@echo 'serve-docs: serve the documentation'

build-docker:
	docker compose build

up-docker:
	docker compose up -d

clean-docker:
	docker compose down -v

remove-docker:
	docker compose down

purge-docker:
	docker compose down -v --rmi all

stop-docker:
	docker compose stop

restart-docker:
	docker compose restart

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