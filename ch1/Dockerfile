# Build the application from source
FROM golang:1.23.4 AS build-stage
WORKDIR /app

# Get dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build sources
COPY cmd/*.go ./
RUN CGO_ENABLED=0 go build -o /kubernetes-playgrounds

# Deploy the application binary into a lean image
FROM scratch

# Copy executable from build stage
WORKDIR /
COPY --from=build-stage /kubernetes-playgrounds /kubernetes-playgrounds

EXPOSE 8080

ENTRYPOINT ["/kubernetes-playgrounds"]
