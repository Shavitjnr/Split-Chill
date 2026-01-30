# Build backend binary file
FROM golang:1.25.5-alpine3.23 AS be-builder
ARG RELEASE_BUILD
ARG BUILD_PIPELINE
ARG BUILD_UNIXTIME
ARG BUILD_DATE
ARG CHECK_3RD_API
ARG SKIP_TESTS
ENV RELEASE_BUILD=$RELEASE_BUILD
ENV BUILD_PIPELINE=$BUILD_PIPELINE
ENV BUILD_UNIXTIME=$BUILD_UNIXTIME
ENV BUILD_DATE=$BUILD_DATE
ENV CHECK_3RD_API=$CHECK_3RD_API
ENV SKIP_TESTS=$SKIP_TESTS
WORKDIR /go/src/github.com/Shavitjnr/splitchill-ai
COPY . .
RUN docker/backend-build-pre-setup.sh
RUN apk add git gcc g++ libc-dev
RUN ./build.sh backend

# Build frontend files
FROM --platform=$BUILDPLATFORM node:24.12.0-alpine3.23 AS fe-builder
ARG RELEASE_BUILD
ARG BUILD_PIPELINE
ARG BUILD_UNIXTIME
ARG BUILD_DATE
ENV RELEASE_BUILD=$RELEASE_BUILD
ENV BUILD_PIPELINE=$BUILD_PIPELINE
ENV BUILD_UNIXTIME=$BUILD_UNIXTIME
ENV BUILD_DATE=$BUILD_DATE
WORKDIR /go/src/github.com/Shavitjnr/splitchill-ai
COPY . .
RUN docker/frontend-build-pre-setup.sh
RUN apk add git
RUN ./build.sh frontend

# Package docker image
FROM alpine:3.23.2
LABEL maintainer="Shavitjnr <shavitjnr@example.com>"
RUN addgroup -S -g 1000 splitchill-ai && adduser -S -G splitchill-ai -u 1000 splitchill-ai
RUN apk --no-cache add tzdata
COPY docker/docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh
RUN mkdir -p /splitchill-ai && chown 1000:1000 /splitchill-ai \
  && mkdir -p /splitchill-ai/data && chown 1000:1000 /splitchill-ai/data \
  && mkdir -p /splitchill-ai/log && chown 1000:1000 /splitchill-ai/log \
  && mkdir -p /splitchill-ai/storage && chown 1000:1000 /splitchill-ai/storage
WORKDIR /splitchill-ai
COPY --from=be-builder --chown=1000:1000 /go/src/github.com/Shavitjnr/splitchill-ai/splitchill-ai /splitchill-ai/splitchill-ai
COPY --from=fe-builder --chown=1000:1000 /go/src/github.com/Shavitjnr/splitchill-ai/dist /splitchill-ai/public
COPY --chown=1000:1000 conf /splitchill-ai/conf
COPY --chown=1000:1000 templates /splitchill-ai/templates
COPY --chown=1000:1000 LICENSE /splitchill-ai/LICENSE
USER 1000:1000
EXPOSE 8080
ENTRYPOINT ["/docker-entrypoint.sh"]
