run: build
	@./bin/goui

build:
	@go build -o bin/goui

css:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch