FROM golang:1.24-bookworm AS builder
WORKDIR /app
COPY ./ /app
COPY ./.env /app/.env
RUN go build -o /notesapi

FROM debian:bookworm-slim
WORKDIR /
COPY --from=builder /notesapi /notesapi
EXPOSE 5000
ENTRYPOINT ["/notesapi"]