run: build
	@./bin/app

build:
	@templ generate
	@go build -tags=prod -o bin/app .
