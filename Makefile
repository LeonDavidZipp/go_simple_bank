startapp:
	docker-compose up

createdb:
	docker-compose exec db createdb -U exampleuser simple_bank

dropdb:
	docker-compose exec db dropdb -U exampleuser simple_bank

sqlc:
	sqlc generate

# A tour about the syntax of go
tour:
	tour

.PHONY: startapp createdb dropdb sqlc