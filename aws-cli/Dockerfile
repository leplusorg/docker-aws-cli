FROM public.ecr.aws/amazonlinux/amazonlinux:2023.9.20251020.0-minimal@sha256:ae11eea8ce5dec8d49d3e939791bf22579ccd8aa802ddaefb8297a5d1889e9bf

HEALTHCHECK NONE

ENTRYPOINT []

ARG USER_NAME=default
ARG USER_HOME=/home/default
ARG USER_ID=1000

SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

RUN dnf upgrade -y \
    && dnf install -y \
    awscli-2-2.30.4-1.amzn2023.0.1 \
    bash-5.2.15-1.amzn2023.0.2 \
    bzip2-1.0.8-6.amzn2023.0.2 \
    curl-minimal-8.11.1-4.amzn2023.0.1 \
    diffutils-3.8-1.amzn2023.0.2 \
    findutils-1:4.8.0-2.amzn2023.0.2 \
    git-2.50.1-1.amzn2023.0.1 \
    golang-1.24.7-1.amzn2023.0.1 \
    grep-3.8-1.amzn2023.0.4 \
    gzip-1.12-1.amzn2023.0.1 \
    iputils-20210202-2.amzn2023.0.4 \
    jq-1.7.1-51.amzn2023 \
    libxml2-2.10.4-1.amzn2023.0.13 \
    make-1:4.3-5.amzn2023.0.2 \
    openssh-8.7p1-8.amzn2023.0.15 \
    patch-2.7.6-14.amzn2023.0.2 \
    pcre2-tools-10.40-1.amzn2023.0.3 \
    pwgen-2.08-11.amzn2023 \
    python3-3.9.23-1.amzn2023.0.3 \
    python3-pip-21.3.1-2.amzn2023.0.14 \
    rsync-3.4.0-1.amzn2023.0.2 \
    sed-4.8-7.amzn2023.0.2 \
    tar-2:1.34-1.amzn2023.0.4 \
    unzip-6.0-57.amzn2023.0.2 \
    wget-1.21.3-1.amzn2023.0.4 \
    xz-5.2.5-9.amzn2023.0.2 \
    zip-3.0-28.amzn2023.0.2 \
    && dnf clean all \
    && rm -rf /var/cache/yum

# FIXME @TL workaround until Amazon Linux gets a pipx package
# FIXME @TL --break-system-packages not supported until Python 3.11.
RUN pip install --no-cache-dir --upgrade pipx==1.8.0

# FIXME @TL remarshal requires python 3.10
RUN pipx ensurepath --global \
    && pipx install --global \
    cfn-lint==1.40.2 \
    cfn-policy-validator==0.0.36 \
    checkov==3.2.487 \
    yamlpath==3.8.2

# FIXME @TL workaround until Amazon Linux gets an yq package
RUN GOBIN=/usr/local/bin go install github.com/mikefarah/yq/v4@v4.47.1

RUN chmod 777 /opt \
    && adduser --home-dir "${USER_HOME}" --uid "${USER_ID}" "${USER_NAME}"

USER "${USER_NAME}"

ENV HOME="${USER_HOME}"

WORKDIR /opt
