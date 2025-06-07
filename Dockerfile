FROM golang:1.24

COPY . .

CMD ["go","run","main.go"]

EXPOSE 8081
