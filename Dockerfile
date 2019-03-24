FROM golang as builder
COPY . /usr/src/httpstaticd
WORKDIR /usr/src/httpstaticd
RUN go get ./... && \
    CGO_ENABLED=0 GOOS=linux go build -o httpstaticd cmd/httpstaticd.go

FROM scratch
COPY --from=builder /usr/src/httpstaticd/httpstaticd /usr/local/bin/httpstaticd
CMD ["/usr/local/bin/httpstaticd"]
