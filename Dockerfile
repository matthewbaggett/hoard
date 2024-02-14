FROM golang:1.21.0 AS builder
ARG VERSION=0.0.1
ARG BUILD_DATE=never
WORKDIR /build

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  hoard

COPY go.* /build
RUN go mod download && \
    go mod verify
COPY ./pkg/ /build/pkg
COPY ./*.go /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags "-X common.version=$VERSION -X common.build_date=$BUILD_DATE" \
      -o datapond \
        datapond.go
RUN chmod +x datapond


FROM scratch AS datapond

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /build/datapond .

USER hoard:hoard
CMD ["./datapond"]
ENTRYPOINT ["./datapond"]
