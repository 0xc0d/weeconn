## Build Stage
#FROM golang:1.15 AS Builder
#
#ENV GO111MODULE=on
#
#WORKDIR $GOPATH/src
#COPY . ./github.com/0xc0d/weeconn
#
#WORKDIR $GOPATH/src/github.com/0xc0d/weeconn
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weeconn
#
#RUN mv weeconn /

# Run Stage
FROM scratch

WORKDIR /

COPY  /weeconn .

EXPOSE 8080

ENTRYPOINT ["./weeconn"]