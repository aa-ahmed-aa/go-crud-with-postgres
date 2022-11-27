# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download -x

COPY . .

RUN go build -o /wallester

EXPOSE 3000

CMD [ "/wallester" ]
