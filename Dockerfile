FROM gocd/gocd-server:v21.2.0

ARG UID

USER root

RUN apk --no-cache add shadow && \
    usermod -u ${UID} go

USER go
