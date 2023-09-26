FROM golang:latest

WORKDIR /app

COPY . .

EXPOSE 3001

RUN go run github.com/steebchen/prisma-client-go db push

CMD go run .