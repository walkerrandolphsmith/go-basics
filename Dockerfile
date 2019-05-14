FROM golang:1.8

RUN mkdir $GOPATH/src/github.com
RUN mkdir $GOPATH/src/github.com/walkerrandolphsmith
RUN mkdir $GOPATH/src/github.com/walkerrandolphsmith/go-playground

WORKDIR /$GOPATH/src/github.com/walkerrandolphsmith/go-playground

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]