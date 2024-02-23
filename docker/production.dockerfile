FROM golang:1.20-bullseye AS builder

ENV GO113MODULE on
WORKDIR /usr/src/app
COPY . .
RUN go build -o /app main.go

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app /app

CMD ["/app"]