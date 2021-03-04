FROM harbor.online.njtech.edu.cn/library/golang:lumanke AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o app .

FROM harbor.jinhun.moe/library/alpine AS runer

WORKDIR /run
COPY --from=builder /build/app .
RUN chmod +x app

ENTRYPOINT ["/run/app"]
