FROM gcr.io/distroless/base-debian10
COPY twitter-cleaner /
ENTRYPOINT ["/twitter-cleaner"]
