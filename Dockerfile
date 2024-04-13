FROM golang:alpine
RUN apk add --no-cache git
WORKDIR /app
COPY go.* .
RUN go mod tidy
COPY . .

CMD ["go", "run", "server.go"]
