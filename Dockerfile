FROM alpine:3.9.6 as dev
WORKDIR /usr/src/app
COPY bin/emaily /usr/local/bin/emaily
COPY templates /usr/local/bin/templates
EXPOSE 8081
HEALTHCHECK --interval=5m --timeout=30s --start-period=5s --retries=10 \
    CMD curl -f http://localhost:8081/check || exit 1
CMD /usr/local/bin/emaily
