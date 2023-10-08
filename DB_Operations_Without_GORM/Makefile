postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres12 createdb --usernames=root --owner=root bankpocs

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankpoc?sslmode=disable" -verbose up

migratedown:
	smigrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankpoc?sslmode=disable" -verbose down

sqlc:
	sqlc generate







.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans redis