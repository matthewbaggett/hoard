FROM golang:1.21.0 AS builder
ARG VERSION=0.0.1
ARG BUILD_DATE=never
WORKDIR /build
COPY go.mod /build
RUN go mod download
RUN go mod verify
#COPY ./pkg/ /build/pkg
#COPY ./*.go /build
COPY . /build
RUN ls -lah /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags "-X common.version=$VERSION -X common.build_date=$BUILD_DATE" \
      -o datapond \
        datapond.go
RUN chmod +x datapond

FROM scratch AS datapond

#COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=builder /etc/passwd /etc/passwd
#COPY --from=builder /etc/group /etc/group

COPY --from=builder /build/datapond .
CMD ["./datapond"]
ENTRYPOINT ["./datapond"]
