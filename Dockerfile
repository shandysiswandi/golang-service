FROM golang:alpine AS builder
ARG TAGGED
LABEL maintainer="SHANDY SISWANDI <shandysiswandi@gmail.com>"
LABEL tagged=$TAGGED
WORKDIR /build
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOGC=off GOOS=linux GOARCH=amd64 go build -a -ldflags="-w -s" -o ./binary ./cmd/echo/main.go

FROM alpine:3.13
LABEL maintainer="SHANDY SISWANDI <shandysiswandi@gmail.com>"
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /build/binary /app/binary
COPY --from=builder /build/.env /app/.env
WORKDIR /app
ENTRYPOINT ["/app/binary"]