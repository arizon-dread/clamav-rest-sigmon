FROM docker.io/golang:1.26-alpine AS build
LABEL MAINTAINER=github.com/arizon-dread
ARG TARGETOS
ARG TARGETARCH
WORKDIR /usr/local/go/src/github.com/arizon-dread/clamav-rest-sigmon
COPY . .

RUN apk update && apk add --no-cache git
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o /usr/local/bin/sigmon/ ./...


FROM dhi.io/alpine-base:3.23 AS final
WORKDIR /go/bin
ARG VERSION
ENV GENERAL_VERSION=${VERSION}
#RUN apk add --no-cache libc6-compat musl-dev
COPY --from=build /usr/local/bin/sigmon/ /go/bin/
EXPOSE 9001
ENTRYPOINT [ "/go/bin/sigmon" ]
