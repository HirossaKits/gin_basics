## About this repository

This repository is for self-learning go backend.

## Set GO PATH

```shell
go env GOPATH
vi ~/.bash_profile
```

and write

```
export PATH=$PATH:$(go env GOPATH)/bin
```

## Set test flags on VS Code

```json
{
  "go.formatTool": "goimports",
  "go.useLanguageServer": true,
  "go.testFlags": ["-v", "-race"]
}
```

## Create DB

https://hub.docker.com/_/postgres

```shell
docker run --name postgres12 -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
```

### Explanation about command

docker run  
--name [container_name]  
-e [environment_variable]  
-p [host_ports:container_ports]  
-d  
[image]:[tag]

### Go into docker container as root user

```shell
docker exec -it postgres12 psql -U root bin/bash
```

### Watch logs

```shell
docker logs
```

### Get IP of container

```shell
docker inspect [ID] | grep IPAddress
```

### Search process using specified port

```shell
lsof -i:5434 -P
```

### Communication between containers

❌ I'm in trouble
https://qiita.com/suo-takefumi/items/bb42f6bd17707de354b9

### Install libraries

```shell
brew install golang-migrate
```

### Stop Container

```shell
docker stop [container name or ID]
```

### Check if container has stopped

```shell
docker ps -a
```

### Restart container

```shell
docker start [container name or ID]
```

### Create and drop DB

```shell
docker exec -it [container name or ID] /bin/sh
createdb --username=root --owner=root [DB Name]
psql [DB Name]
dropdb [DB Name]
\q
exit
```

### Watch history

```shell
history | grep "docker run"
```

## Drop docker container

```shell
docker rm [container name or ID]
```

## Migrate DB

```shell
migrate -path db/migration -database "postgres://root:secret@localhost:5434/simple_bank?sslmode=disable" -verbose up
```

## Install sqlc

```shell
brew install sqlc
```

## Create sqlc setting file

```shell
sqlc init
```

## Generate model

```shell
sqlc generate
```

## Install testify

```shell
go get github.com/stretchr/testify
```

## Make random int

```go
import (
  "fmt"
  "time"
  "math/rand"
)
rand.Seed(time.Now().Unix())
fmt.Println(rand.Intn(10))
```

## Read Input

```go
import "fmt"
var input string
fmt.Scan(&input)
```

## Install Gin

```shell
go get -u github.com/gin-gonic/gin
```

## Install mockgen

```shell
go get github.com/golang/mock/mockgen@v1.6.0
go install github.com/golang/mock/mockgen@v1.6.0
```

## Add path to mockgen

```shell
vi ~/.zshrc
```

And add below line

```shell
export PATH=$PATH:~/go/bin
```

## Apply settings to shell

```shell
source ~/.zshrc
```

## Confirm place of mockgen

```shell
which mockgen
```

## Create mock DB

```shell
mockgen -destination db/mock/store.go gin_basics/db/sqlc Store
```

## Rename package of mock DB

```shell
mockgen -package mockdb -destination db/mock/store.go gin_basics/db/sqlc Store
```
