FROM golang:alpine

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /opt
COPY src/* /opt/fib-service/

WORKDIR /opt/fib-service
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o ./bin/fib-service
# Make sure the binary is executable
RUN chmod +x ./bin/fib-service

# Set the entrypoint to run the binary
ENTRYPOINT ["/opt/fib-service/bin/fib-service"]
