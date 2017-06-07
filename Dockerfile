FROM golang:1.7.5 as builder

WORKDIR /go/src/github.com/redhatcop/gows

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gows .

FROM scratch

WORKDIR /opt

COPY --from=builder /go/src/github.com/redhatcop/gows/gows /bin/gows

EXPOSE 8080

CMD ["/bin/gows"]
