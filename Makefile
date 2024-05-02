startapp:
	docker-compose up

createdb:
	docker-compose exec db createdb -U exampleuser simple_bank

dropdb:
	docker-compose exec db dropdb -U exampleuser simple_bank

sqlc:
	sqlc generate

.PHONY: startapp createdb dropdb sqlc