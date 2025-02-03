FROM golang:alpine AS builder

RUN apk add --no-cache git

WORKDIR /build
RUN git clone https://github.com/linux4life798/webdav-server.git . && \
    go build

FROM alpine:latest

COPY --from=builder /build/webdav-server /usr/local/bin/

RUN mkdir /webdav && \
    chmod 777 /webdav

# Mount your share directory to /webdav.
VOLUME /webdav

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/webdav-server"]
CMD ["-port", "8080", "-dir", "/webdav"]
