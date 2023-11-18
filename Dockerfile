#syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o go-example

FROM scratch AS runner
WORKDIR /app
COPY --from=builder /app/go-example ./go-example
EXPOSE 8080
CMD [ "/app/go-example" ]
