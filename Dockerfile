FROM golang:1.15.11-buster
RUN mkdir /usr/local/go/src/helloGinAndGorm
ADD . /usr/local/go/src/helloGinAndGorm
WORKDIR /usr/local/go/src/helloGinAndGorm

# Install Git
RUN apt-get install git -y

# Install package for server.go
RUN go get -u github.com/gin-gonic/gin
RUN go get -u gorm.io/gorm

# Build
RUN go build server.go
CMD ["/usr/local/go/src/helloGinAndGorm/server"]