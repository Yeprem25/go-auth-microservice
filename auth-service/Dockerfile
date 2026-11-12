FROM golang:1.25-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /out/api-gateway ./api-gateway/cmd

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /out/api-gateway /api-gateway

ENV GIN_MODE=release
EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/api-gateway"]
