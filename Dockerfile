FROM golang:1.22 AS builder
ENV PROJECT_PATH=/app/server
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY . ${PROJECT_PATH}
WORKDIR ${PROJECT_PATH}
RUN go build cmd/server/main.go

FROM golang:alpine
WORKDIR /app/cmd/server
COPY --from=builder /app/server/main .
EXPOSE 30002
CMD ["./main"]
