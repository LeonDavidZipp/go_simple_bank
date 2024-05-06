start:
	docker-compose up

startapp:
	docker-compose up app

startdb:
	docker-compose up db \
	&& make createdb \
	&& make migrateup

createdb:
	docker-compose exec db createdb -U exampleuser simple_bank

migrateup:
	docker-compose run --rm app migrate -path ./db/migrations -database "postgresql://exampleuser:test1234@db:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	docker-compose run --rm app migrate -path ./db/migrations -database "postgresql://exampleuser:test1234@db:5432/simple_bank?sslmode=disable" -verbose down

restartdb: dropdb createdb migrateup

dropdb:
	docker-compose exec db dropdb -U exampleuser simple_bank

runcmd:
	docker-compose exec db psql -U exampleuser -d simple_bank -c "$(cmd)"

sqlc:
	docker-compose run --rm app sh -c "sqlc generate"

test:
	docker-compose run --rm app sh -c "go test -v -cover ./..."

get:
	docker-compose run --rm app go get -u $(pkg)

imagebuild:
	docker-compose build

imagerebuild:
	docker-compose build --no-cache

# A tour about the syntax of go
tour:
	tour

.PHONY: start startapp startdb createdb migrateup migratedown dropdb sqlc test tour
