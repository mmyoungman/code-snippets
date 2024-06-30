include .env

install:
	@npm install

run: build
	@./bin/app

build:
	@./node_modules/.bin/tailwindcss -i views/layouts/base.css -o public/styles.css
	@templ generate
	@go build -tags=prod -o bin/app .

watch:
	${MAKE} -j4 watch-tailwindcss watch-templ watch-go

watch-tailwindcss:
	@./node_modules/.bin/tailwindcss -i views/layouts/base.css -o public/styles.css --watch

watch-templ:
	@templ generate -watch -proxy=http://localhost:${LISTEN_PORT}

watch-go:
	@air