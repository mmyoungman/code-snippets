include .env

install:
# @MarkFix specify versions?
# updating templ beyond .833 breaks - not this but related: https://github.com/a-h/templ/issues/1146
	@go install github.com/a-h/templ/cmd/templ@v0.3.833
	@go install github.com/air-verse/air@latest
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install github.com/go-jet/jet/v2/cmd/jet@latest
	@go mod download
	@${MAKE} goose cmd="up"
	@${MAKE} jet
	@go install golang.org/x/vuln/cmd/govulncheck@latest

run: build
	@./bin/main

kc:
	@docker-compose up keycloak

build:
	@${MAKE} jet
	@templ generate
	@go fmt ./...
	@go build -o bin/main .

buildProd:
	@templ generate
	@go build -tags=prod -o bin/main .

clean:
	@rm -rf bin
	@rm -rf tmp
	@find . -name '*_templ.go' -delete
	@find . -name '*_templ.txt' -delete

watch:
# air handles running templ generate for us - see cmd in .air.toml
	@air -c .air.toml

# Options for hot reloading - search for @hotreload to see how it works
watchProxy:
	@${MAKE} -j4 watchTemplProxy watchGoProxy
watchTemplProxy:
	@templ generate -watch -proxy=${PUBLIC_HOST}:${PUBLIC_PORT} -v
watchGoProxy:
	PROXY_URL=http://127.0.0.1:7331 air -c .air.proxy.toml

goose:
# make migration cmd="create new_migration_name sql"
# make migration cmd="up"
# see all goose commands here: https://github.com/pressly/goose?tab=readme-ov-file#usage
	@goose -v -s -dir "${MIGRATIONS_PATH}" sqlite3 "${SQLITE3_PATH}" ${cmd}

jet:
	@jet -source=sqlite -dsn="${SQLITE3_PATH}" -path="${JETS_PATH}"

format:
	@go fmt ./...

govulncheck:
	@govulncheck ./...