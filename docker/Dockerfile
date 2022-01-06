#build stage
FROM golang:1.16.2-alpine AS builder
WORKDIR /go/src/log-kit
COPY . .
RUN apk add --no-cache git
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go install  -v ./...

#final stage
FROM alpine:3.11
RUN apk --no-cache add ca-certificates
WORKDIR /opt/log-kit
COPY --from=builder /go/bin/* .
COPY --from=builder /go/src/log-kit/*.yaml .

