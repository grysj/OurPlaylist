postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest
createdb:
	docker exec -it postgres15 createdb --username=root --owner=root ourplaylist_db

dropdb:
	docker exec -it postgres15 dropdb ourplaylist_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost/ourplaylist_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost/ourplaylist_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go


.PHONY: postgres createdb dropdb migrateup migratedown sqlc