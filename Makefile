createpg:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
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
	migrate -path dbMigration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
up: deletepg createpg createdb migrateup sqlc
# https://stackoverflow.com/questions/2145590/what-is-the-purpose-of-phony-in-a-makefile
.PHONY: createpg deletepg createdb dropdb migrateup migratedown sqlc up