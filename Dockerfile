FROM golang:1.22-rc as builder

WORKDIR /builds

ADD * ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fah-exporter

FROM gcr.io/distroless/static-debian11

COPY --from=builder /builds/fah-exporter .

COPY --from=builder /builds/config.yaml .

CMD ["./fah-exporter"]
