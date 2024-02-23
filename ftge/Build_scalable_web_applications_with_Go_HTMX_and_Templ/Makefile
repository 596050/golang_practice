APP_NAME=dreampicai

run: build
	@./bin/$(APP_NAME)

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

build:
	tailwindcss -i view/css/app.css -o public/styles.css
	@templ generate view
	@go build -o bin/$(APP_NAME) main.go

up: ## Database migration up
	@go run cmd/migrate/main.go up

drop:
	@go run cmd/drop/main.go up

down: ## Database migration down
	@go run cmd/migrate/main.go down

migration: ## Migrations against the database
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

gen:
	@go run cmd/generate/main.go

seed:
	@go run cmd/seed/main.go