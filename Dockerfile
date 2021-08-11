FROM golang:1.16 as builder

RUN mkdir /build
WORKDIR /build
COPY . /build
ENV CGO_ENABLED=0
RUN go mod vendor
RUN go build -o api

FROM scratch
COPY --from=builder /build/api /api
EXPOSE 8080

CMD ["/api"]


