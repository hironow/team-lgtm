FROM golang:1.12 as builder

ENV GO111MODULE=on
WORKDIR /go/src/github.com/hironow/team-lgtm

# Warm up dependency cache
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build binary
COPY . .
RUN cd frontend && CGO_ENABLED=0 GOOS=linux go build -v -o frontend

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/hironow/team-lgtm/frontend/frontend /frontend

# Run the web service on container startup.
CMD ["/frontend"]