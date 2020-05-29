FROM alpine:3.9.6 as dev
WORKDIR /usr/src/app
COPY bin/emaily /usr/local/bin/emaily
COPY templates /usr/local/bin/templates
EXPOSE 8081
CMD /usr/local/bin/emaily