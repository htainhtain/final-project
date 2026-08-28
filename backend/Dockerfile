FROM golang:alpine AS stage-build

WORKDIR /usr/local/app

COPY . /usr/local/app

RUN go build && \
    go test ./... && \
    rm -rf \
    /root/.cache/go-build \
    /go/pkg/mod

#run time stage
FROM alpine:latest

WORKDIR /usr/local/app

EXPOSE 8080

COPY --from=stage-build /usr/local/app/server ./

RUN adduser -D appuser && chown appuser /usr/local/app

USER appuser

ENTRYPOINT [ "./server" ]