FROM golang:1.21-alpine3.18

RUN apk update && apk add git curl

ENV GOPATH /go

RUN mkdir /go/app
COPY . /go/app

WORKDIR /go/app

# sql
RUN go get -u github.com/go-sql-driver/mysql &&\
    go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# SQLBoiler
RUN go install github.com/volatiletech/sqlboiler/v4@latest &&\
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest &&\
    go get github.com/volatiletech/sqlboiler/v4 &&\ 
    go get github.com/volatiletech/null/v8

RUN go get -u github.com/gocarina/gocsv &&\
    go install github.com/joho/godotenv/cmd/godotenv@latest
