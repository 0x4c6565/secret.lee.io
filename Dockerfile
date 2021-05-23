FROM golang:1.16-alpine AS build
WORKDIR /build
COPY . .
RUN go build -o secret.lee.io

FROM alpine
WORKDIR /app
COPY --from=build /build/secret.lee.io .
COPY static static
ENTRYPOINT ["/app/secret.lee.io"]
