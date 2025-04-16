# syntax=docker/dockerfile:1

FROM golang:1.24-alpine
WORKDIR /app

RUN apk add --no-cache --update gcc g++ curl

# Install air to a shared location
RUN go install github.com/air-verse/air@v1.61.7
RUN cp /go/bin/air /usr/local/bin/
RUN chmod +x /usr/local/bin/air

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
