# # # # Start from a Debian-based image with Go installed
# # # FROM golang:latest as build

# # # # Set the current working directory inside the container
# # # WORKDIR /app

# # # # Copy the Go modules and files needed for building the app
# # # COPY go.mod go.sum ./

# # # # Download and install Go dependencies
# # # RUN go mod download

# # # # Copy the rest of the application source code
# # # COPY . .

# # # # Build the Go application
# # # RUN go build -o main .

# # # # Expose port 3000 to the outside world
# # # # EXPOSE 3000

# # # # # Command to run the executable
# # # # CMD ["./main"]

# # # FROM gcr.io/distroless/static-debian11
# # # COPY --from=build /app/main .
# # # EXPOSE 3000
# # # CMD ["./main"]


# FROM golang:latest as build

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .
# RUN go build -o main .

# FROM gcr.io/distroless/static-debian11
# COPY --from=build /app/main .
# EXPOSE 3000
# CMD ["./main"]

# Build Stage
FROM golang:latest as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final Stage
FROM gcr.io/distroless/static-debian11

COPY --from=build /app/main /

EXPOSE 3000

CMD ["/main"]
