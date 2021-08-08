# Golang binglog

## Setup database
binlog format must be 'row' In pkg/mariadb.go change with your connection data
```
cfg.Addr = fmt.Sprintf("%s:%d", "127.0.0.1", 3306) //host,port
cfg.User = "root"
cfg.Password = "root"
```
## Run Database Server
```
docker-compose up
```

## Create Table Schema
Create database with sql
```
create table user
(
  id int auto_increment primary key,
  name varchar(40) null,
  status enum("active","deleted") DEFAULT "active",
  created timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
) engine=InnoDB;
```
Set bin log on in query
```
SET sql_log_bin = 1;
```

## Run
```
go run src/main.go
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
