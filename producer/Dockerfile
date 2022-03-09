FROM php:8.1-apache
LABEL maintainer "Luis Rodriguez"

RUN apt-get update && apt-get install --no-install-recommends -y \
    ca-certificates \
    build-essential \
    software-properties-common \
    librdkafka-dev

RUN pecl config-set php_ini "/usr/local/etc/php/php.ini"
RUN pecl install rdkafka-6.0.1
