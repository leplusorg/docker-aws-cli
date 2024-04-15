FROM amazon/aws-cli:2.15.38@sha256:0fadd0fc0cd5c415d103cc8187ae8744a66b085da33cc56d2d2ceccc306e2a97

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
