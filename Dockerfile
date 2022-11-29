FROM golang:1.19-alpine AS build-env

ARG MONIKER
ENV MONIKER moniker

ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

COPY . /go/delivery/defund

COPY ./devtools/entrypoint.sh /

RUN chmod +x /entrypoint.sh

WORKDIR /go/delivery/defund

RUN make install

RUN defundd init $MONIKER

EXPOSE 26656 26657 1317 9090

CMD ["/bin/bash", "/entrypoint.sh"]