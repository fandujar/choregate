FROM node:latest as frontend-builder
WORKDIR /app
COPY . .
RUN cd choregate-ui && yarn install && yarn build

FROM golang:latest as backend-builder
WORKDIR /app
COPY --from=frontend-builder /app /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./cmd/choregate ./cmd/choregate 

FROM ubuntu:latest as linux-setup
RUN useradd -m choregate

FROM scratch as final
WORKDIR /app
COPY --from=backend-builder /app/cmd/choregate/choregate /choregate
COPY --from=linux-setup /etc/passwd /etc/passwd
USER choregate
ENTRYPOINT ["/choregate"]
EXPOSE 8080