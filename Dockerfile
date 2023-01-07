FROM golang:1.19.4-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY . .

RUN make build

CMD ["/app/main"]

EXPOSE 1102