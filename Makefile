postgres: 
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres:16-alpine

remove-postgres:
	docker stop postgres16
	docker rm postgres16

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root fasunga

dropdb:
	docker exec -it postgres16 dropdb fasunga

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fasunga?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fasunga?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres remove-postgres createdb dropdb migrateup migratedown sqlc test server