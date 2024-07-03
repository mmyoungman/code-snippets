include .env

install:
	@npm install

run: build
	@./bin/main

kc:
	@docker-compose up keycloak

build:
	@./node_modules/.bin/tailwindcss -i views/layouts/base.css -o public/styles.css
	@templ generate
	@go build -o bin/main .

buildProd:
	@./node_modules/.bin/tailwindcss -i views/layouts/base.css -o public/styles.css --minify
	@templ generate
	@go build -tags=prod -o bin/main .

clean:
	@rm -rf bin
	@rm -rf tmp
	@find . -name '*_templ.go' -delete
	@find . -name '*_templ.txt' -delete
	@rm public/styles.css

watch:
# air handles running templ generate for us - see cmd in .air.toml
	@${MAKE} -j4 watchTailwindcss watchGo

watchTailwindcss:
	@./node_modules/.bin/tailwindcss -i views/layouts/base.css -o public/styles.css --watch

watchGo:
	@air -c .air.toml

# Below options are to get templ browser hot reloading working with auth - search for @hotreload to see how it works
watchProxy:
	@${MAKE} -j4 watchTailwindcss watchTemplProxy watchGoProxy

watchTemplProxy:
	@templ generate -watch -proxy=${PUBLIC_HOST}:${PUBLIC_PORT} -v

watchGoProxy:
	TEMPL_WATCH_PROXY_URL=://127.0.0.1:7331 air -c .air.proxy.toml
