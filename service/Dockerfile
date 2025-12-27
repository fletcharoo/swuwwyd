FROM golang:latest AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./service ./service
RUN go build -o /app/main ./service

FROM debian:latest
RUN mkdir -p /pb_data
COPY --from=build /app/main /app/main
ENTRYPOINT ["/app/main", "serve", "--http=0.0.0.0:8080"]
