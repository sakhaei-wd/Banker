createdb:
	docker exec -it postgres-0 createdb --username=postgres --owner=postgres postgres

dropdb:
	docker exec -it postgres-0 dropdb postgres

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
