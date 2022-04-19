FROM golang:alpine

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go build -o applikasiku .

CMD ["./applikasiku"]