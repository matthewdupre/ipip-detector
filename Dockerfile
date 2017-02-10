FROM alpine:3.4
MAINTAINER Matt Dupre <matt@projectcalico.org>

COPY bin/ipip-detector /


CMD "/ipip-detector"