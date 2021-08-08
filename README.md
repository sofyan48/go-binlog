# Golang binglog

## Configuration
### Environment
```
$cp .env.example .env
```
### Run Compose
```
$docker-compose up
```
### Migration
```
go run src/main.go run db:migrate up
```
## Run
```
go run src/main.go run
```
than execute sql in your app
```
INSERT Into user (`id`,`name`) VALUE (1,"Jack");
UPDATE user SET name="Jonh" WHERE id=1;
DELETE FROM user WHERE id=1;

```

report if finish and successfully
```
User 1 is created with name Jack
User 1 is updated from name Jack to name Jonh
User 1 is deleted with name Jonh
```

## Example Architecture
![Example Architecture](https://github.com/sofyan48/go-binlog/blob/master/resource/ex_architecture.png)
