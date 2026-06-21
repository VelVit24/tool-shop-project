include .env
migrate-up:
	migrate \
	-path database/migrations \
	-database "$(DATABASE_URL)" \
	up ${VERSION}

migrate-down:
	migrate \
	-path database/migrations \
	-database "$(DATABASE_URL)" \
	down ${VERSION}

migrate-create:
	migrate create -ext sql -dir database/migrations -seq ${NAME}

migrate-force:
	migrate \
		-path database/migrations \
		-database "$(DATABASE_URL)" \
		force ${VERSION}