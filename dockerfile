#Stage 1 : Build Go binary

FROM golang:1.23-alpine AS builder 
WORKDIR /myapp
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o cookbook-app

#Stage 2 : Small runtime image

FROM alpine:latest
WORKDIR /myapp
COPY --from=builder /myapp/cookbook-app .
EXPOSE 8080
CMD ["./cookbook-app"]
