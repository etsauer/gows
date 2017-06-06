FROM golang:1.7.5

WORKDIR /go/src/github.com/redhatcop/gows

RUN go get -d -v golang.org/x/net/http
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gows .

FROM scratch

WORKDIR /opt

ADD --from=0 /go/src/github.com/redhatcop/gows/gows /bin/gows

EXPOSE 8080

CMD ["/bin/gows"]
