FROM golang:latest

WORKDIR $GOPATH/src/bug-log/go/docker
COPY . $GOPATH/src/bug-log/go/docker
RUN go build .

EXPOSE 8000
ENTRYPOINT [ "./docker" ]