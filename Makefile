init-db:
	@docker-compose up -d postgres
	@docker-compose exec postgres psql -U myuser auth_user_db -f /docker-entrypoint-initdb.d/database.sql
