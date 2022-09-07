# Build the manager binary
FROM golang:1.19 as builder
ARG cert_location=/usr/local/share/ca-certificates

RUN openssl s_client -servername proxy.golang.org -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/proxy-golang-org.crt
RUN ls /usr/local/share/ca-certificates


RUN update-ca-certificates
WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go source
COPY main.go main.go
COPY open/ open/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o aperta main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/aperta .
USER nonroot:nonroot

ENTRYPOINT ["/aperta"]
