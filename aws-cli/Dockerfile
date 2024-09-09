FROM amazon/aws-cli:2.17.46@sha256:8b62ffc6dc7eff2ec25640c5d65a65e5dd5acecdfebea2c619ab74ab452bea47

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

# @TL FIXME: --break-system-packages not supported until Python 3.11
# hadolint ignore=DL3013
RUN pip3 install --no-cache-dir --upgrade cfn-policy-validator

RUN mkdir -p /opt \
    && chmod 777 /opt \
    && adduser --home-dir "${USER_HOME}" --uid "${USER_ID}" "${USER_NAME}"

ENV HOME "${USER_HOME}"

USER "${USER_NAME}"

WORKDIR /opt

CMD ["/bin/bash"]
