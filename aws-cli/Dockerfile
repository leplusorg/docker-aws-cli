FROM amazon/aws-cli:2.15.38@sha256:0fadd0fc0cd5c415d103cc8187ae8744a66b085da33cc56d2d2ceccc306e2a97

HEALTHCHECK NONE

ENTRYPOINT []

ARG USER_NAME=default
ARG USER_HOME=/home/default
ARG USER_ID=1000

# hadolint ignore=DL3033
RUN yum update -y \
    && yum install -y bash curl git jq make python3-pip tar unzip xmlstarlet zip \
    && yum clean all \
    && rm -rf /var/cache/yum

# hadolint ignore=DL3013
RUN pip3 install --no-cache-dir cfn-policy-validator

RUN adduser \
  --home-dir "${USER_HOME}" \
  --uid "${USER_ID}" \
  "${USER_NAME}"

ENV HOME "${USER_HOME}"

USER "${USER_NAME}"

WORKDIR /opt

CMD ["/bin/bash"]
