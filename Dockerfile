FROM golang:1.21.1-alpine3.18

COPY ./src /app

WORKDIR /app