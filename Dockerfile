FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV GOPROXY=https://goproxy.cn,direct
RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#https://mirrors.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app/stuoj

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN ls
RUN go run ./dev/tools/clean_generated.go -y
RUN go generate ./...
RUN go build -ldflags="-s -w" -o ./stuoj main.go

FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/stuoj/stuoj /app/stuoj

CMD ["./stuoj"]

EXPOSE 14514/tcp