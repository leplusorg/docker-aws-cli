FROM amazon/aws-cli:2.5.8

# hadolint ignore=DL3033
RUN yum update -y \
    && yum install -y bash curl git golang jq make tar unzip xmlstarlet zip \
    && yum clean all \
    && rm -rf /var/cache/yum \
    && set -o pipefail
    
ENV GOVERSION="$(\go version | \cut -d ' ' -f 3)"
ENV GO_INSTALL_DIR=/usr/lib/golang
ENV GOROOT="${GO_INSTALL_DIR}"
ENV GOPATH=/root/golang
ENV PATH="${GOROOT}/bin:${GOPATH}/bin:${PATH}"
ENV GO111MODULE=on
ENV GOSUMDB=off

RUN go get -u github.com/raviqqe/liche

WORKDIR /opt

ENTRYPOINT []

CMD ["/bin/bash"]
