postgres:
	docker run --name basketball --network bdb-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=basketball -d postgres:alpine3.14

postgres-wo-network:
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

build_docker:
	docker build -t basketball_db:latest .

run_docker:
	docker run --name b_db --network bdb-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:basketball@basketball:5432/b_db?sslmode=disable" basketball_db 

.PHONY: postgres createdb dropdb migrateup migratedown sqlc build_docker run_docker