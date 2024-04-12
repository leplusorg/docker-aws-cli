FROM amazon/aws-cli:2.15.37@sha256:0ad4fcf33c409b843d7cd67c7b944936b1ad4ab96ad8d84baf44c8009f6cccb1

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
