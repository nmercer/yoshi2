
FROM golang:1.14 as builder

RUN apt-get update

WORKDIR $GOPATH/src/github.com/nmercer/yoshi2/services/server
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /app .

EXPOSE 50051
ENTRYPOINT ["/app"]
