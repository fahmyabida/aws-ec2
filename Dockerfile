FROM golang:alpine as tempat-buildnya

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go build -o applikasiku .

FROM alpine
RUN mkdir /jadinya
COPY --from=tempat-buildnya /app/applikasiku  /jadinya/applikasiku

CMD ["./jadinya/applikasiku"]