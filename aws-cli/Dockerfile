FROM amazon/aws-cli:2.15.36@sha256:9e4f6d7ef854f0fa42c161042ddc51994ef654f5f9f66a9b41ba6d5896e9dfc0

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
