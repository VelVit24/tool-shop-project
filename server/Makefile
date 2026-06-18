migrate-up:
	migrate \
	-path database/migrations \
	-database "$(DATABASE_URL)" \
	up

migrate-down:
	migrate \
	-path database/migrations \
	-database "$(DATABASE_URL)" \
	down