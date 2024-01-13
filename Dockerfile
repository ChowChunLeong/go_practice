FROM golang:latest

WORKDIR /app

COPY . .

RUN go get

EXPOSE 8080

CMD ["go", "run", "."]