table = $(table)
name = $(name)
url=mysql://db-user:db-password@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local
version=$(version)
migration-up :
	migrate -database "$(url)" -path ./migrations/ up $(version)
	
migration-down :
	migrate -database "$(url)" -path ./migrations/ down $(version)
	
migration-create:
	migrate create -ext sql -dir ./migrations/ -seq $(name)

migration-force:
	migrate -database "$(url)" -path ./migrations/ force $(version)

migration-version:
	migrate -database "$(url)" -path ./migrations/ version

run :
	gow run .

