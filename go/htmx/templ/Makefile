include .env

install:
# @MarkFix specify versions?
# sqlc v1.28 onwards causes this https://github.com/sqlc-dev/sqlc/issues/3852
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27
# updating templ beyond .833 breaks - not this but related: https://github.com/a-h/templ/issues/1146
	@go install github.com/a-h/templ/cmd/templ@v0.3.833
	@go install github.com/air-verse/air@latest
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@mkdir -p ${MIGRATIONS_PATH}
	@go mod download
	@${MAKE} goose cmd="up"

build:
	@sqlc generate
	@templ generate
	@go fmt ./...
	@go build -o bin/main .

run: build
	@./bin/main

buildProd:
	@sqlc generate
	@templ generate
	@go build -tags=prod -o bin/main-prod .

runProd:
	@./bin/main-prod

clean:
	@rm -rf bin
	@rm -rf tmp
	@find . -name '*_templ.go' -delete
	@find . -name '*_templ.txt' -delete
	@go mod tidy

kc:
	@docker-compose up keycloak

watch:
# air handles running templ generate for us - see cmd in .air.toml
	@air -c .air.toml

# Options for hot reloading - search for @hotreload to see how it works
watchHotreload:
	@${MAKE} -j4 watchTemplProxy watchGoProxy
watchTemplHotreload:
	@templ generate -watch -proxy=${PUBLIC_HOST}:${PUBLIC_PORT} -v
watchGoHotreload:
	PROXY_URL=http://127.0.0.1:7331 air -c .air.proxy.toml

goose:
# make migration cmd="create new_migration_name sql"
# make migration cmd="up"
# see all goose commands here: https://github.com/pressly/goose?tab=readme-ov-file#usage
	@goose -v -s -dir "${MIGRATIONS_PATH}" sqlite3 "${SQLITE3_PATH}" ${cmd}

format:
	@go fmt ./...

govulncheck:
	@govulncheck ./...