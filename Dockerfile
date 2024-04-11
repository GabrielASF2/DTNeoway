FROM golang:1.22-bookworm

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download

# Install the postgres driver here
RUN go install github.com/lib/pq

RUN go build -o buyers ./

CMD ["./buyers"]
