FROM golang:1.19.0-alpine3.16 AS builder 

RUN mkdir /app

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

COPY .env .

RUN go build -o main .

FROM alpine

RUN mkdir /apps

WORKDIR /apps

COPY .env /apps/

COPY --from=builder /app/main .

CMD ["/apps/main"]