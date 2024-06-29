include .env

install:
	@npm install

run: build
	@./bin/app

build:
	@templ generate
	@go build -tags=prod -o bin/app .
	@./node-modules/.bin/tailwindcss -i views/css/app.css -o public/styles.css

watch:
	${MAKE} -j4 watch-go watch-templ watch-tailwindcss

watch-go:
	@air

watch-templ:
	@templ generate -watch -proxy=http://localhost${LISTEN_ADDR}

watch-tailwindcss:
	@./node_modules/.bin/tailwindcss -i views/css/app.css -o public/styles.css --watch