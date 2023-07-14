# build stage
FROM golang:1.19-alpine as build

RUN mkdir /App

ADD . /App

WORKDIR /App

RUN go build -o main .

CMD ["/App/main"]