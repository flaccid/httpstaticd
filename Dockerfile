FROM golang as builder
COPY . /go/src/github.com/flaccid/httpstaticd
WORKDIR /go/src/github.com/flaccid/httpstaticd
RUN go get ./... && \
    CGO_ENABLED=0 GOOS=linux go build -o httpstaticd cmd/httpstaticd/httpstaticd.go

FROM scratch
COPY --from=builder /go/src/github.com/flaccid/httpstaticd/httpstaticd /httpstaticd
WORKDIR /
ENTRYPOINT ["./httpstaticd"]
