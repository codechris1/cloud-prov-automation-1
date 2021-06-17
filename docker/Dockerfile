FROM golang:latest

WORKDIR /go/src/app
COPY go/src/* .

#RUN go get -d -v ./...
RUN go install -v ./hello.go

CMD ["go", "run", "hello.go"]