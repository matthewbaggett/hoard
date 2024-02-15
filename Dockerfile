FROM golang:1.21.0 AS builder
ARG VERSION=0.0.1
ARG BUILD_DATE=never
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
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

FROM builder as builder-datapond
COPY ./datapond /build/datapond
RUN go build -o datapond datapond/datapond.go
RUN chmod +x datapond

FROM scratch AS datapond
COPY --from=builder-datapond /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder-datapond /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder-datapond /etc/passwd /etc/passwd
COPY --from=builder-datapond /etc/group /etc/group
COPY --from=builder-datapond /build/datapond .
USER hoard:hoard
CMD ["./datapond"]
ENTRYPOINT ["./datapond"]

FROM builder as builder-datalake
COPY ./datalake /build/datalake
RUN go build -o datalake datalake/datalake.go
RUN chmod +x datapond

FROM scratch AS datalake
COPY --from=builder-datalake /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder-datalake /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder-datalake /etc/passwd /etc/passwd
COPY --from=builder-datalake /etc/group /etc/group
COPY --from=builder-datalake /build/datalake .
USER hoard:hoard
CMD ["./datalake"]
ENTRYPOINT ["./datalake"]

