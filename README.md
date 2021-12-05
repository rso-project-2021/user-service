# user-service
![Build](https://github.com/rso-project-2021/user-service/actions/workflows/build.yml/badge.svg)
![Deploy](https://github.com/rso-project-2021/user-service/actions/workflows/deploy.yml/badge.svg)  
Microservice used for working with user data.

## Environment file
In root of your local repository add `config.json` file.
```
{
    "db_driver" : "postgres",
    "db_source": "postgres://root:secret@localhost:5432/user_service?sslmode=disable",
    "server_address": "0.0.0.0:8080",
    "gin_mode": "debug"
}
```

## Setup database
1. Run `docker pull postgres:alpine` to download [postgres image](https://hub.docker.com/_/postgres).
2. Run `make postgres` to run postgres image inside of container.
3. Run `make createdb` to create postgres database.
4. Run `make migrateup` to add "users" table.
5. Run `go mod tidy` to clean golang package dependecies.
6. Test project with command `make test`.
7. Run service with `go run .`.
8. Use [PostMan](https://www.postman.com/) to send query to `http://localhost:8080/v1/users/`.

## Seed database
Populate database with some users. You can run this query in [TablePlus](https://tableplus.com/).
```sql
INSERT INTO users("username", "password", "email")
VALUES 	('Mario', 'passgancipass', 'ganci@gmail.com'),
	('Johnny', 'john123', 'johnny@gmail.com'),
	('Donald Trump', 'trumpyboy', 'trump@hotmail.com'),
	('Obama', 'barackog44', 'obama@gmail.com');
```

## Things to implement
- [x] CRUD operations
- [x] Database migrations
- [x] CRUD unit tests
- [x] Makefile
- [x] Health checks
- [x] Docker file
- [x] CI github actions
- [x] Dockerhub
- [x] AWS account
- [x] Kubernetes cluster in AWS
- [x] Metrics ([Prometheus in Go](https://prometheus.io/docs/guides/go-application/))
- [x] CD github actions
- [ ] Config server (dynamic configuration)
- [ ] API unit tests
