# Stage 1 - building app
FROM golang:1.17-alpine3.15 AS BUILD

ARG APP_NAME=${APP_NAME:-app}
ARG ENTRYPOINT='deploy/entrypoint.sh'
ARG SOPS_VERSION
WORKDIR /src/

# download modules in separated layer, to speed up rebuild by utilising Docker layer caching system
COPY go.mod go.sum /src/
# NOTE: build error may occur due to temporary unavailability of some packages sources
# Wait and build again is usually a good solution
RUN set -xe &&\
    go mod download

COPY . /src/

RUN set -xe &&\
    . /etc/os-release && _OS=$(uname -s | tr '[:upper:]' '[:lower:]') &&\
    apk --update --no-cache add gnupg jq &&\
    wget -qO /usr/local/bin/sops $(wget -qO- https://api.github.com/repos/mozilla/sops/releases/latest | jq -r '.assets[] | select(.name | endswith("'${_OS}'")) | .browser_download_url') || exit 1 &&\
    chmod +x /usr/local/bin/sops /src/${ENTRYPOINT} &&\
    sops --version || wget -qO /usr/local/bin/sops https://github.com/mozilla/sops/releases/download/v${SOPS_VERSION##v}/sops-v${SOPS_VERSION##v}.linux || exit 1 &&\
    chmod +x /usr/local/bin/sops /src/${ENTRYPOINT} &&\
    wget -qO- https://dl.dropboxusercontent.com/s/noe19dbp3mxyzgf/appconfig.asc | gpg --import - &&\
    sops -d --input-type yaml /src/config.enc > /src/config.toml || exit 1

RUN set -xe &&\
    go build -o bin/${APP_NAME} ./cmd/app/main.go

# Stage 2 - serving app
FROM alpine:3.15

ARG APP_NAME=${APP_NAME:-app}
ARG PORT_EXPOSE=${PORT_EXPOSE}
ARG ENTRYPOINT
ENV APP_NAME=${APP_NAME}\
    ENTRYPOINT=${ENTRYPOINT}
WORKDIR /app/

COPY --from=BUILD /src/${ENTRYPOINT} /src/bin/${APP_NAME} /src/*.yaml /src/*.toml /src/*.enc /app/

EXPOSE ${PORT_EXPOSE}

ENTRYPOINT ["/app/entrypoint.sh"]
