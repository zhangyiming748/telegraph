FROM golang:1.22.5-bookworm
LABEL authors="zen"
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/root/go/bin
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod vendor
ENTRYPOINT ["go","run","/app/main.go"]