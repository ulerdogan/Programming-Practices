postgres:
	docker run --name basketball -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=basketball -d postgres:alpine3.14

createdb:
	docker exec -it basketball createdb --username=root --owner=root b_db

dropdb:
	docker exec -it basketball dropdb b_db

migrateup:
	migrate -path db/migration -database "postgresql://root:basketball@localhost:5432/b_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:basketball@localhost:5432/b_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres