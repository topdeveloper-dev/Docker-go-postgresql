FROM golang:1.17-buster as builder

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD ["/app/main"]