FROM arigaio/atlas:0.36.2-alpine

RUN apk update && apk add --no-cache ca-certificates curl && rm -rf /var/cache/apk/*

COPY migrations /migrations
