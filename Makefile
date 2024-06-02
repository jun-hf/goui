run: build
	@./bin/goui

build:
	@go build -o bin/goui

css:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch

templWatch:
	templ generate --watch --proxy=http://localhost:3000