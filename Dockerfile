FROM golang:1.19.1-alpine3.16 AS Builder

WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=Builder /app/main .
COPY app.env .
EXPOSE 8080
CMD ["/app/main"]