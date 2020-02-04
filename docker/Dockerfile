FROM golang:1.13.7-alpine

ADD ./src /src
WORKDIR /src

RUN go get

CMD go run app.go
