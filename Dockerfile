FROM harbor.jinhun.moe/library/alpine AS runer

COPY /build/app .
RUN chmod +x app

ENTRYPOINT ["./app"]
