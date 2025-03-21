FROM node:latest AS frontend-build
WORKDIR /app/frontend
COPY ./frontend/ /app/frontend/
RUN npm i
RUN npm run build

FROM golang:1.24-bookworm AS builder
WORKDIR /app
COPY ./ /app
COPY ./.env /app/.env
RUN GIN_MODE=release CGO_ENABLED=0 GOOS=linux go build -o /notesapi

FROM alpine:latest
WORKDIR /
COPY --from=frontend-build /app/frontend/dist/ /frontend/dist/
COPY --from=builder /app/.env /.env
COPY --from=builder /notesapi /notesapi
EXPOSE 5000
ENTRYPOINT ["/notesapi"]