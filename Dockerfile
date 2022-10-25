# syntax=docker/dockerfile:1
# docker build --tag docker-viper .
# docker build --tag adrianwudev/docker-viper .

FROM golang:1.18-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# COPY *.go ./
COPY . ./
RUN go build -o /docker-viper
CMD [ "/docker-viper" ]