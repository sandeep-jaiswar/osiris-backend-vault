# Use the official Golang image to create a build artifact.
FROM golang:1.20 AS builder

# Copy local code to the container image.
WORKDIR /app
COPY . .

# Build the Go app with static linking.
RUN CGO_ENABLED=0 GOOS=linux go build -tags netgo -ldflags '-extldflags "-static"' -o /app/osiris-backend-vault

# Use a minimal base image to package the compiled binary.
FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/osiris-backend-vault /app/osiris-backend-vault

# Run the web service on container startup.
CMD ["/app/osiris-backend-vault"]
