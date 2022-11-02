FROM golang:alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/system

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git

# Configure these values
# ENV GOPRIVATE="github.com/<your private org>/*"
# RUN git config --global url."ssh://git@github.com/<your private org>".insteadOf "https://github.com/<your private org>" 

# Build Go binary
COPY Makefile go.mod go.sum ./
RUN make init && go mod download 
COPY . .
RUN make proto tidy build

# Deployment container
FROM scratch

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/system/system /system
ENTRYPOINT ["/system"]
CMD []
