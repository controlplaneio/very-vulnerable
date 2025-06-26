FROM golang:1.25rc1-bookworm as builder
WORKDIR /
COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod download
RUN go build -o /very-vulnerable

#https://hub.docker.com/layers/library/debian/trixie-20240904-slim/images/sha256-d80214935426b02c162971a3213e1a14312f01a632c68f0ef09f27dcd6f33267
# Contains a collection of Critical, High and Med Vulns
FROM debian:trixie-20240904-slim as runner 

RUN mkdir /download

COPY github_token.txt /download/github_token.txt
COPY id_ed25519 /download/id_ed25519

COPY --from=builder /very-vulnerable /very-vulnerable

EXPOSE 8080

# Hardcoded credentials in ENV variables
ENV GITHUB_TOKEN="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"

USER root

ENTRYPOINT [ "/very-vulnerable" ]
