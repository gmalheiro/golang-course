DB_URL=postgresql://root:root@localhost:5432/simple-bank?sslmode=disable


startdb:
	docker-compose up -d 

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple-bank

dropdb:
	docker exec -it postgres dropdb simple-bank

test:
	go test -v -cover ./...

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test ./... -coverprofile=/tmp/coverage.out -covermode=atomic -race -count=1 -shuffle=on


.PHONY: startdb createdb dropdb migrateup migratedown sqlc 
