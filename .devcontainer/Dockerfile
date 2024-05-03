FROM --platform=linux/amd64 golang:latest

RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - \
    && echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ xenial main" > /etc/apt/sources.list.d/migrate.list \
    && apt-get update \
    && apt-get install -y migrate \
    && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# go introduction
RUN go install golang.org/x/website/tour@latest

ENV dbDriver="postgres"
ENV dbSource="postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
# ENV dbSource="postgresql://root:secret@host.docker.internal:5432/simple_bank?sslmode=disable"
