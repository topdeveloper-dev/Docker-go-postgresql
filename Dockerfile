FROM golang:1.17-buster as builder

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod /app

COPY go.sum /app

RUN go mod download

COPY . /app

EXPOSE 8084

CMD ["./main"]