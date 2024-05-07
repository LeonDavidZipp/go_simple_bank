########################################################################################
#                                                                                      #
#    Base                                                                              #
#                                                                                      #
########################################################################################

FROM golang:1.22.2-alpine3.19 as base

LABEL maintainer="lzipp"

WORKDIR /app
COPY ./db ./db
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

EXPOSE 8000


########################################################################################
#                                                                                      #
#    Development                                                                       #
#                                                                                      #
########################################################################################

FROM base as dev

ENV dbDriver="postgres"
ENV dbSource="postgresql://exampleuser:test1234@db:5432/simple_bank?sslmode=disable"
ENV serverAddress="0.0.0.0:8000"

COPY ./sqlc.json ./

RUN go install golang.org/x/website/tour@latest \
    && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest


########################################################################################
#                                                                                      #
#    Production                                                                        #
#                                                                                      #
########################################################################################

# FROM base as prod
