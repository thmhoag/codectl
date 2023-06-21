# Start from goreleaser so we can pull it's entrypoint
# and binary
FROM goreleaser/goreleaser as goreleaser

# Move to golang to ensure we're on the desired version
FROM golang:1.20.5-alpine as build-sdk

RUN apk add --no-cache \
    bash \
    curl \
    git

COPY --from=goreleaser /entrypoint.sh /entrypoint.sh
COPY --from=goreleaser /usr/bin/goreleaser /usr/bin/goreleaser

ENTRYPOINT ["/entrypoint.sh"]
CMD [ "-h" ]
