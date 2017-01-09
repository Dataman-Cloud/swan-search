FROM golang:1.6.3-alpine

COPY . /go/src/github.com/Dataman-Cloud/swan-search
WORKDIR /go/src/github.com/Dataman-Cloud/swan-search
RUN go build -v -o bin/search

EXPOSE 9888

ENTRYPOINT ["/go/src/github.com/Dataman-Cloud/swan-search/bin/search"]
