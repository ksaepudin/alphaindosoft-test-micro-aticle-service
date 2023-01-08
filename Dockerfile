FROM golang:1.19.4-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY . .

RUN go build -o ./main cmd/server/rest/main.go

CMD ["/app/main"]

EXPOSE 1102