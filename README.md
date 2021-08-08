# Golang binglog

## Example Architecture
![Example Architecture](https://github.com/sofyan48/go-binlog/blob/master/resource/ex_architecture.png)


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

