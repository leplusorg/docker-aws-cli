FROM amazon/aws-cli:2.24.21@sha256:a4a0bd1569f12e4428811bab066ca84b72864c3b901699f50644678792bae9ba

HEALTHCHECK NONE

ARG USER_NAME=default
ARG USER_HOME=/home/default
ARG USER_ID=1000

# hadolint ignore=DL3033
RUN yum update -y \
    && yum install -y bash curl git jq make python3-pip tar unzip xmlstarlet zip \
    && yum clean all \
    && rm -rf /var/cache/yum

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# @TL FIXME: --break-system-packages not supported until Python 3.11
RUN pip3 install --no-cache-dir --upgrade \
    cfn-policy-validator==0.0.35 \
    && msg="$(pip list --outdated | grep -i -e cfn-policy-validator || true)" \
    && if [ -n "${msg}" ]; then \
       >&2 echo "ERROR: outdated: ${msg}" ; \
       exit 1 ; \
    fi

RUN mkdir -p /opt \
    && chmod 777 /opt \
    && adduser --home-dir "${USER_HOME}" --uid "${USER_ID}" "${USER_NAME}"

ENV HOME="${USER_HOME}"

USER "${USER_NAME}"

WORKDIR /opt

CMD ["/bin/bash"]

ENTRYPOINT []
