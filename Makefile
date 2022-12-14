createpg:
	sudo docker run --name postgres12 --network bank-nw -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
	sleep 30
deletepg:
	sudo docker stop postgres12
	sudo docker rm postgres12
createdb:
	sudo docker exec postgres12 createdb --username=root --owner=root simple_bank
dropdb:
	sudo docker exec postgres12 dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
migrateup1:
# here 1 at end is version, check schema_migrations table for version info
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown1:
# here 1 at end is version,  check schema_migrations table for version info
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/luckyparakh/goBank/db/sqlc Store
up: deletepg createpg createdb migrateup sqlc

server:
	go run main.go
# https://stackoverflow.com/questions/2145590/what-is-the-purpose-of-phony-in-a-makefile
.PHONY: createpg deletepg createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc up server mock