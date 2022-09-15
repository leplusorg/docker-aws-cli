FROM amazon/aws-cli:2.7.32

# hadolint ignore=DL3033
RUN yum update -y \
    && yum install -y bash curl git jq make tar unzip xmlstarlet zip \
    && yum clean all \
    && rm -rf /var/cache/yum

WORKDIR /opt

ENTRYPOINT []

CMD ["/bin/bash"]
