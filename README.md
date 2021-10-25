# user-service
Microservice used for retrieving and manipulating user data.

### Docker commands for postgres
- Run postgres image in container: `docker run --name image_name -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres` 
- Stop postgres container: `docker stop image_name`
- Opens postgres in terminal: `docker exec -t image_name psql -U root`

### Environment file
In your local repository add `app.env` file.
```
DB_DRIVER=postgres
DB_SOURCE=postgres://root:password@localhost:5432/root?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
GIN_MODE=debug
```
### Database initialization
Sample postgresql query to instantiate database schema.
```sql
DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users (
    "user_id"        BIGSERIAL PRIMARY KEY,
    "username"       	VARCHAR(40) NOT NULL,
    "password"       	VARCHAR(40) NOT NULL,
    "email" 			VARCHAR(256),
    "created_at"      	TIMESTAMP NOT NULL DEFAULT(now())
);

INSERT INTO users("username", "password", "email")
VALUES 	('Mario', 'passgancipass', 'ganci@gmail.com'),
		('Johnny', 'john123', 'johnny@gmail.com'),
		('Donald Trump', 'trumpyboy', 'trump@hotmail.com'),
		('Obama', 'barackog44', 'obama@gmail.com')
RETURNING *;
```
