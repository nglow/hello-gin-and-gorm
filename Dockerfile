FROM golang:1.15.11-buster
RUN mkdir /app
ADD . /app
WORKDIR /app

# Build
RUN go build server.go
CMD ["/app/server"]