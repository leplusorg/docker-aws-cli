FROM amazon/aws-cli:2.15.35@sha256:9d1582256677c240d1f735b4ff4cc407c0596dccffa713fc405186794fab8139

# hadolint ignore=DL3033
RUN yum update -y \
    && yum install -y bash curl git jq make python3-pip tar unzip xmlstarlet zip \
    && yum clean all \
    && rm -rf /var/cache/yum

# hadolint ignore=DL3013
RUN pip3 install --no-cache-dir cfn-policy-validator

WORKDIR /opt

ENTRYPOINT []

CMD ["/bin/bash"]
