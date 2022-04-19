FROM golang:alpine

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN apk update
RUN apk add git

RUN go build -o applikasiku .

CMD ["./applikasiku"]