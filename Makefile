start:
	docker-compose up

startapp:
	docker-compose up app

startdb:
	docker-compose up db

createdb:
	docker-compose exec db createdb -U exampleuser simple_bank

migrateup:
	docker-compose run --rm app migrate -path ./db/migrations -database "postgresql://exampleuser:test1234@db:5432/simple_bank?sslmode=disable" -verbose up

dropdb:
	docker-compose exec db dropdb -U exampleuser simple_bank

sqlc:
	docker-compose run --rm app sh -c "sqlc generate"

testaccount:
	docker-compose run --rm app sh -c "go test ./db/sqlc -run TestCreateAccount"

# A tour about the syntax of go
tour:
	tour

.PHONY: startapp createdb dropdb sqlc