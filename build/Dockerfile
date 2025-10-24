FROM alpine:3.6

USER nobody

ADD templates/*.yaml /usr/local/bin/templates/
ADD build/_output/bin/gitea-operator /usr/local/bin/gitea-operator

RUN apk add --no-cache git=2.47.3-r0 && \
    mkdir -p /usr/bin/git_2.47-r0 && \
    cp /usr/bin/git /usr/bin/git_2.47-r0/git