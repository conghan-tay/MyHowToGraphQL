postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root hacker_noon

dropdb:
	docker exec -it postgres12 dropdb hacker_noon

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/hacker_noon?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/hacker_noon?sslmode=disable" -verbose down

sqlc:
	sqlc generate

postgresconsole:
	docker exec -it postgres12 psql -U root -d hacker_noon

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test