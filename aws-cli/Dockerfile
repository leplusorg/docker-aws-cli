FROM amazon/aws-cli:2.4.23

ENV NODE_ENV production

# hadolint ignore=DL3033
RUN yum update -y \
    && yum install -y bash curl git jq make tar unzip zip \
    && yum clean all \
    && rm -rf /var/cache/yum

WORKDIR /opt

ENTRYPOINT []

CMD ["/bin/bash"]
