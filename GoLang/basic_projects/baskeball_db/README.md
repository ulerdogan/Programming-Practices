following practices for [the course](https://www.youtube.com/playlist?list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE)

### requirements

for the db table preperation: [dbdiagram](https://dbdiagram.io)
- draw the tables -which is in dbml format-
- then export to postgresql -in sql format- 

for the db table management gui: [tableplus](https://tableplus.com/)
- run sql scripts to add tables


### do's

- preparing and using the docker container

```
docker run --name techmaster -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=techmaster -d postgres:alpine3.14

docker exec -it techmaster psql -U root

docker logs techmaster

docker stop techmaster

docker start techmaster

docker ps

docker rm techmaster
```

- manage db in the container

``` shell
// runs in the shell
docker exec -it techmaster /bin/sh

// creates db in the shell
createdb --username=root --owner=root simple_bank

// connect
psql simple_bank

// deletes
dropdb simple_bank

// do directly from docker query
docker exec -it techmaster psql -U root simple_bank

```

- download ```golang-migrate``` tool for db migrations

    ```brew install golang-migrate```

- creates migration files

    ```migrate create -ext sql -dir db/migration -seq init_schema```

- prepare up file from dbdiagram export, and down with 'drop's

- run migrations by ```make migrateup``` or ```make migratedown```

- download sqlc ```brew install sqlc```

- generate sqlc files by ```sqlc init```, add+edit tags, create folders, write query, then ```sqlc generate```