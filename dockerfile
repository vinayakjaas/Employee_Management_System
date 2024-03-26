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
