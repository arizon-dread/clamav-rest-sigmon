FROM docker.io/golang:1.26-alpine AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /usr/local/go/src/github.com/arizon-dread/clamav-rest-sigmon
COPY . .

RUN apk update && apk add --no-cache git
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o /usr/local/bin/sigmon/ ./...


FROM dhi.io/alpine-base:3.23 AS final
WORKDIR /go/bin
ARG TARGETOS
ARG TARGETARCH
ARG VERSION
ENV GENERAL_VERSION=${VERSION}
LABEL MAINTAINER=github.com/arizon-dread \
  org.opencontainers.image.version=${VERSION} \
  org.opencontainers.image.os=${TARGETOS} \
  org.opencontainers.image.arch=${TARGETARCH}
#RUN apk add --no-cache libc6-compat musl-dev
COPY --from=build /usr/local/bin/sigmon/ /go/bin/
EXPOSE 9001
ENTRYPOINT [ "/go/bin/sigmon" ]
