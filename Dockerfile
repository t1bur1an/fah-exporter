FROM golang:1.22-rc as builder

WORKDIR /builds

ADD * ./

RUN go build -o fah-exporter

FROM golang:1.22-rc

WORKDIR /app

COPY --from=builder /builds/fah-exporter /app/fah-exporter

COPY --from=builder /builds/config.yaml /app/config.yaml

CMD ["/app/fah-exporter"]
