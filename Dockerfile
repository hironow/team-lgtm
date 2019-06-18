FROM golang:1.12

ENV GO111MODULE=on
WORKDIR /go/src/github.com/hironow/team-lgtm

# Warm up dependency cache
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Run test
CMD ["go", "test", "./..."]
