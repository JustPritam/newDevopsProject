FROM golang:1.24

WORKDIR /app

COPY . .

CMD ["go","run","main.go"]

