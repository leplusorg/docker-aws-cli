FROM amazon/aws-cli:2.15.39@sha256:0e58eddb945a5496c5e4370b37676b06f7ac54d40d5fb98c93df2ea749f0ad11

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


RUN mkdir /opt && chmod 777 /opt

ENV HOME "${USER_HOME}"

USER "${USER_NAME}"

WORKDIR /opt

CMD ["/bin/bash"]
