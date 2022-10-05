FROM golang:1.19.1-alpine3.16 AS Builder

WORKDIR /app
COPY . .
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
RUN go build -o main main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=Builder /app/main .
COPY --from=Builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
RUN chmod +x start.sh
COPY ./db/migration ./migration
EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]