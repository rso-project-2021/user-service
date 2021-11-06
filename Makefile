postgres:
	docker run --name mypostgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it mypostgres createdb --username=root --owner=root user_service

dropdb:
	docker exec -it mypostgres dropdb user_service

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/user_service?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/user_service?sslmode=disable" -verbose down

test: 
	go test -cover ./...

.PHONY: postgres, createdb, dropdb, migrateup, migratedown, test