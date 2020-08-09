FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git make

WORKDIR /app
COPY . .

RUN make vendor
RUN make build

FROM alpine:3
RUN apk update \
    && apk add --no-cache curl \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates 2>/dev/null || true


COPY --from=builder /app/bin/url_short /url_short

CMD ["/url_short"]

ENV PORT=8080
EXPOSE $PORT
