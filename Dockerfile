FROM golang

COPY . /go/src/github.com/sikupe/corona-rules

WORKDIR /go/src/github.com/sikupe/corona-rules

RUN go get ./

RUN go build

EXPOSE 8080

CMD ["./corona-rules"]