FROM golang:1.22

RUN apt-get update \
     && apt-get clean --quiet --yes \
     && apt-get autoremove --quiet --yes \
     && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./olo_service/cmd/main.go

CMD ["./main"]