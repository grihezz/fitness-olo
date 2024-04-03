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

ENV CONFIG_PATH auth_service/resourse/config.yml


EXPOSE 5500

CMD ["./main"]