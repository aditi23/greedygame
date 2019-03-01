# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/aditi23/greedygame/auction

RUN go install github.com/aditi23/greedygame/auction

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/auction

# Document that the service listens on port 8080.
EXPOSE 9999