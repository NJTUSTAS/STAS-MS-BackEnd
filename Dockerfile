FROM harbor.jinhun.moe/library/alpine:latest

COPY app .
RUN chmod +x app

ENTRYPOINT ["./app"]
