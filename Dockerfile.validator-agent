# Build the validator-agent binary
FROM docker.io/golang:1.25 AS builder
WORKDIR /build
COPY validator-agent validator-agent
RUN cd validator-agent && go build -o validator-agent .

FROM debian:latest
WORKDIR /app

# Install necessary runtime dependencies
RUN apt update && apt install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the validator-agent binary from builder
COPY --from=builder /build/validator-agent/validator-agent /app

# Copy the config file
COPY validator-agent/config.testnet.yaml /app/config.yaml
#COPY validator-agent/config.mainnet.yaml /app/config.yaml

# Set the entrypoint
CMD ["/app/validator-agent"]
