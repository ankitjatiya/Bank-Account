postgres:
	docker run --name postgres18 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16rc1-alpine3.18

createdb:
	docker exec -it postgres18 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres18 dropdb bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go Bank-Account/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock

