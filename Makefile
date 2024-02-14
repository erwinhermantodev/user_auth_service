init-db:
	@docker-compose up -d postgres
	@docker-compose exec postgres psql -U myuser auth_user_db -f /docker-entrypoint-initdb.d/database.sql

run:
	go run cmd/main.go

init:
	go mod tidy
	go mod vendor

generated: api.yml
	@echo "Generating files..."
	mkdir generated || true
	oapi-codegen --package=api --generate types,server,spec api.yml > generated/api.gen.go

mock-generate: 
	cd ./repository && mockery --name=UserRepositoryInterface

tests:
	go test -v ./...
