FROM amazon/aws-cli:2.5.8

# hadolint ignore=DL3033
RUN yum update -y \
    && yum install -y bash curl git golang jq make tar unzip xmlstarlet zip \
    && yum clean all \
    && rm -rf /var/cache/yum \
    && export GOVERSION="$(\go version | \cut -d ' ' -f 3)" \
    && export GO_INSTALL_DIR=/usr/lib/golang \
    && export GOROOT="${GO_INSTALL_DIR}" \
    && export GOPATH=/root/golang \
    && export PATH="${GOROOT}/bin:${GOPATH}/bin:${PATH}" \
    && export GO111MODULE=on \
    && export GOSUMDB=off \
    && go get -u github.com/raviqqe/liche

WORKDIR /opt

ENTRYPOINT []

CMD ["/bin/bash"]
