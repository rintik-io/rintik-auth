FROM golang:1.19.2-bullseye AS builder

RUN apt-get install git

WORKDIR /rintik-io/rintik-auth

COPY . .

RUN go mod vendor

FROM golang:1.19.2-bullseye

ENV GOVERSION="go1.19.2" \
    GO111MODULE="auto" \
    GOCACHE=/tmp \
    XDG_CACHE_HOME=/tmp/.parameter

WORKDIR /rintik-io/rintik-auth

COPY --from=builder /rintik-io/rintik-auth ./

USER root

RUN apt-get update && \
    apt-get install -y nano && \
    chmod -R 777 /rintik-io/rintik-auth

EXPOSE 8600

CMD ["/bin/bash","/rintik-io/rintik-auth/scripts/run.sh"]