# Build the manager binary
FROM golang:1.17 as builder

WORKDIR /workspace
ENV GOPROXY="https://goproxy.io"

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
## cache deps before building and copying source so that we don't need to re-download as much
## and so that source changes don't invalidate our downloaded layer
RUN go mod download

## Copy the go source
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GO111MODULE=on go build -a -o demo main.go

FROM amazonlinux
WORKDIR /
COPY --from=builder /workspace/demo .
USER 65532:65532
EXPOSE 4000
ENTRYPOINT [ "/demo" ]
CMD /demo