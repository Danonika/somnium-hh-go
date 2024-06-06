FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive
ENV LANG en_US.UTF-8

# Add base environment
RUN apt-get -qq update \
    && apt-get -qqy --no-install-recommends install \
    curl \
    ca-certificates > /dev/null \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* \
    && update-ca-certificates

ARG GO_FILE
ARG VER

WORKDIR /app

COPY "$GO_FILE" app
COPY Makefile /app/
COPY deploy/ /app/deploy
#COPY scripts/ /app/scripts/
COPY docs/swagger /app/docs/swagger

RUN set -x                                  && \
    echo "${PROJECT_NAME}-${VER}" > version && \
    chmod +x app

ENTRYPOINT ["/app/app"]
