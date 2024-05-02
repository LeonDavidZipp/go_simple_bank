FROM golang:1.22.2-alpine3.19

RUN apk add --no-cache curl
RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apk add --no-cache - 
RUN echo "https://packagecloud.io/golang-migrate/migrate/alpine/v3.7/main" > /etc/apk/repositories
RUN apk update
RUN apk add migrate
