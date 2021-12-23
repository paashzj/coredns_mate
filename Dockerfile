FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o coredns_mate .


FROM ttbb/coredns:nake

LABEL maintainer="shoothzj@gmail.com"

COPY docker-build /opt/sh/coredns/mate

COPY --from=build /opt/sh/compile/pkg/coredns_mate /opt/sh/coredns/mate/coredns_mate

WORKDIR /opt/sh/coredns

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/sh/coredns/mate/scripts/start.sh"]