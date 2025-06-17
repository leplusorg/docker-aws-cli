FROM public.ecr.aws/amazonlinux/amazonlinux:2023.7.20250609.0-minimal

HEALTHCHECK NONE

ARG USER_NAME=default
ARG USER_HOME=/home/default
ARG USER_ID=1000

SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

RUN dnf upgrade -y \
    && dnf install -y --setopt=skip_missing_names_on_install=False \
    aws-cli \
    curl \
    git \
    libxml2 \
    jq \
    make \
    python3 \
    tar \
    unzip \
    zip \
    && dnf clean all \
    && rm -rf /var/cache/yum

# @TL FIXME: cfn-policy-validator version is held back by Python 3.8.
# @TL FIXME: --break-system-packages not supported until Python 3.11.
RUN pip3 install --no-cache-dir --upgrade --break-system-packages \
    cfn-policy-validator==0.0.29 \
    && IFS=$'\n\t' \
    && msg="$(pip3 list --outdated | grep -i -e '^cfn-policy-validator ' || true)" \
    && if [ -n "${msg}" ]; then \
       >&2 echo "ERROR: outdated: ${msg}" ; \
       exit 1 ; \
    fi

RUN chmod 777 /opt \
    && adduser --home-dir "${USER_HOME}" --uid "${USER_ID}" "${USER_NAME}"

USER "${USER_NAME}"

ENV HOME="${USER_HOME}"

WORKDIR /opt

CMD ["/bin/bash"]

ENTRYPOINT []
