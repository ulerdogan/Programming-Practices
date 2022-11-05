docker run -it -d -p 5432:5432 --name gotr-postgre -e POSTGRES_PASSWORD=gotr-pass -d postgres:alpine3.14

docker run -d -p 15672:15672 -p 5672:5672 --name my-rabbit rabbitmq:3-management-alpine

docker run -p 6379:6379 --name some-redis -d redis