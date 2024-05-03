FROM golang:latest

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install golang.org/x/website/tour@latest

ENV dbDriver="postgres"
ENV dbSource="postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"