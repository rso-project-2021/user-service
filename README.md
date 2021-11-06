# user-service
![Build](https://github.com/rso-project-2021/user-service/actions/workflows/ci.yml/badge.svg)  
Microservice used for working with user data.

## Environment file
In root of your local repository add `app.env` file.
```
DB_DRIVER=postgres
DB_SOURCE=postgres://root:secret@localhost:5432/user_service?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
GIN_MODE=debug
```

## Setup database
1. Run `docker pull postgres:alpine` to download [postgres image](https://hub.docker.com/_/postgres).
2. Run `make postgres` to run postgres image inside of container.
3. Run `make createdb` to create postgres database.
4. Run `make migrateup` to add "users" table.
5. Run `go mod tidy` to clean golang package dependecies.
6. Test project with command `make test`.
7. Run service with `go run .`.
8. Use [PostMan](https://www.postman.com/) to send query to `http://localhost:8080/v1/users/1`.

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
- [ ] API unit tests
- [ ] Health checks
- [x] Docker file
- [ ] CI/CD github actions
- [ ] Dockerhub
- [ ] AWS account
