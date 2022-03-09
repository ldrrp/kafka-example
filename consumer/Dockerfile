FROM alpine:latest
LABEL maintainer "Luis Rodriguez"

WORKDIR /opt/consumer/

COPY --from=golang:alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

RUN chmod +x /opt/consumer/entrypoint.sh

EXPOSE 80

CMD /opt/consumer/entrypoint.sh