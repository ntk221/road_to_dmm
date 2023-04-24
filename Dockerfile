# dev, builder
FROM golang:1.16 AS golang
WORKDIR /work/yatter-backend-go

# dev
FROM golang as dev
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# builder
FROM golang AS builder
RUN apt-get update && apt-get -y install build-essential gcc wget
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.33-r0/glibc-2.33-r0.apk && \
    wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.33-r0/glibc-bin-2.33-r0.apk && \
    wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.33-r0/glibc-i18n-2.33-r0.apk
RUN apt-get update && apt-get -y install ca-certificates git tzdata curl
COPY ./ ./
RUN make prepare build-linux

# release
FROM debian:11 AS app
RUN apt-get update && apt-get -y install tzdata
RUN ln -fs /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=builder /work/yatter-backend-go/build/yatter-backend-go-linux-amd64 /usr/local/bin/yatter-backend-go
EXPOSE 8080
ENTRYPOINT ["yatter-backend-go"]