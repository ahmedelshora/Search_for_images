FROM golang:1.24.2-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

# Just run directly without building
CMD ["go", "run", "."]