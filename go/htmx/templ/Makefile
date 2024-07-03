include .env

install:
	@npm install

run: build
#	@docker-compose up keycloak
	@./bin/app

build:
	@./node_modules/.bin/tailwindcss -i views/layouts/base.css -o public/styles.css --minify
	@templ generate
	@go build -tags=prod -o bin/app .

clean:
	@rm -rf bin
	@rm -rf tmp
	@find . -name '*_templ.go' -delete
	@find . -name '*_templ.txt' -delete
	@rm public/styles.css

#keycloak:
#	@docker-compose up keycloak

watch:
#	${MAKE} -j4 watch-tailwindcss watch-templ watch-go
	${MAKE} -j4 watch-tailwindcss watch-go

watch-tailwindcss:
	@./node_modules/.bin/tailwindcss -i views/layouts/base.css -o public/styles.css --watch

#watch-templ:
#	@templ generate -watch -proxy=${PUBLIC_PORT}

watch-go:
	@air