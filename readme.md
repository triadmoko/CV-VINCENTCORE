# Test Teknis VINCENTCORE

## TECH STACK

- Golang
- MySQL
- GORM
- Clean Architecture
- Go Migrate
- Gow auto reload golang

## How to use

### Prerequisite

- Docker
- Docker-compose
- Go-Migrate

### Running Project

```bash
go mod tidy

docker-compose up

make migration-up version=1
```

## API

file documentation api `VINCENTCORE.postman_collection.json` import in postman collection
