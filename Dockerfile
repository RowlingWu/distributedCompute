FROM golang:1.8
RUN mkdir -p /go/src/homework
ADD . /go/src/homework
WORKDIR /go/src/homework
CMD ["go", "run", "main.go"]
