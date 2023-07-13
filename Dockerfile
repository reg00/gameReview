FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /game-review

RUN go build -o main .

FROM alpine AS final

WORKDIR /app
COPY --from=build /app/main ./
COPY --from=build /app/internal/infrastructure/config/config.yaml ./internal/infrastructure/config/

EXPOSE 8080

CMD ["./main"]
