FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

FROM golang:1.21-alpine as runner

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8080

CMD [ "./main" ]