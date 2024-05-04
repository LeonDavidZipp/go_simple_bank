FROM golang:latest

EXPOSE 8000

ARG DEV=false

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
    && if [ $DEV = true ]; \
        go install golang.org/x/website/tour@latest; \
    fi

COPY ./app /app

ENV dbDriver="postgres"
ENV dbSource="postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
