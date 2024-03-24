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

ENV CONFIG_PATH api_gateway/resourse/config.yml

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./api_gateway/cmd/main.go

EXPOSE 5000

CMD ["./main"]