FROM public.ecr.aws/amazonlinux/amazonlinux:2023.7.20250609.0-minimal@sha256:c278930b7f4d5b703962bbcfb7a0ac4c6dc6c318ee02dd8123005d68cd94df17

HEALTHCHECK NONE

ARG USER_NAME=default
ARG USER_HOME=/home/default
ARG USER_ID=1000

SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

RUN dnf upgrade -y \
    && dnf install -y \
    awscli-2-2.23.11-1.amzn2023.0.1 \
    curl-minimal \
    git-2.47.1-1.amzn2023.0.3 \
    gzip-1.12-1.amzn2023.0.1 \
    libxml2-2.10.4-1.amzn2023.0.10 \
    jq-1.7.1-49.amzn2023.0.2 \
    make-1:4.3-5.amzn2023.0.2 \
    python3-3.9.22-1.amzn2023.0.1 \
    tar-2:1.34-1.amzn2023.0.4 \
    unzip-6.0-57.amzn2023.0.2 \
    zip-3.0-28.amzn2023.0.2 \
    && IFS=$'\n\t' \
    && if dnf upgrade -y | grep -v -e '^Nothing to do.' ; then \
       exit 1 ; \
    fi \
    && dnf clean all \
    && rm -rf /var/cache/yum

# @TL FIXME: --break-system-packages not supported until Python 3.11.
RUN pip3 install --no-cache-dir --upgrade \
    cfn-policy-validator==0.0.36 \
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
