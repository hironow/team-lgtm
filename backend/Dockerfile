FROM golang:1.12 as builder

ENV GO111MODULE=on
WORKDIR /go/src/github.com/hironow/team-lgtm

# Warm up dependency cache
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build binary
COPY backend ./backend
RUN cd backend && CGO_ENABLED=0 GOOS=linux go build -v -o backend gql/server/server.go

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine

# Add ca-certificates
# https://github.com/googleapis/google-cloud-go/issues/928
RUN apk --no-cache --update add ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/hironow/team-lgtm/backend/backend /backend

# Run the web service on container startup.
CMD ["/backend"]