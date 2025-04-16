FROM arigaio/atlas:0.32.0-alpine

RUN apk update && apk add --no-cache ca-certificates curl && rm -rf /var/cache/apk/*

COPY migrations /migrations
