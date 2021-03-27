FROM harbor.jinhun.moe/library/alpine:latest

COPY /build/app .
RUN chmod +x app

ENTRYPOINT ["./app"]
