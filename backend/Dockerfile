# syntax=docker/dockerfile:1

# Apply multi-stage build process

# Use standard debian based image to build with stability
FROM golang:1.21.5-bullseye AS build

WORKDIR app

# Copy only files necessary for installing deps for better layer caching
COPY go.* ./

RUN go mod download

EXPOSE 8080

COPY *.go ./

# Run with static build (faster build time)
RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o /go-backend main.go

# Use scratch image for production deployment (minize image size)
FROM scratch

WORKDIR /

COPY --from=build /go-backend /go-backend

# default endpoint used by GIN
EXPOSE 8080

CMD ["/go-backend"]