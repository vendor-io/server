FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .

RUN apk add --no-cache git
RUN apk add --no-cache gcc musl-dev
COPY ./.env /
RUN go build -v -o app .

FROM alpine:latest
WORKDIR /root
COPY ./.env /
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app .
ENTRYPOINT ENV=DEV ./app
EXPOSE 3000