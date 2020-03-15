
############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add git


RUN mkdir /src && mkdir /src/hmdl-user-service
ADD . /src/hmdl-user-service
WORKDIR /src/hmdl-user-service

RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hmdl-user-service .

FROM alpine:latest
RUN apk add --no-cache tzdata
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ENV TZ Asia/Ho_Chi_Minh
ENV VN Asia/Ho_Chi_Minh

WORKDIR /app

COPY --from=builder /src/hmdl-user-service /app/
VOLUME /app/config
EXPOSE 7001

ENTRYPOINT ["./hmdl-user-service"]