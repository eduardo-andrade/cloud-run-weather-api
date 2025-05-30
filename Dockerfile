FROM golang:1.23 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/cloudrun .
COPY .env .
EXPOSE 8080
ENTRYPOINT ["./cloudrun"]
