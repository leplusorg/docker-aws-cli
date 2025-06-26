FROM public.ecr.aws/amazonlinux/amazonlinux:2023.7.20250623.1-minimal@sha256:90ac129125cb499a927ac91d1767808234cf974d26e43557efa0c8ff020c4cad

HEALTHCHECK NONE

ARG USER_NAME=default
ARG USER_HOME=/home/default
ARG USER_ID=1000

SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

RUN dnf upgrade -y \
    && dnf install -y \
    awscli-2-2.23.11-1.amzn2023.0.1 \
    bash-5.2.15-1.amzn2023.0.2 \
    bzip2-1.0.8-6.amzn2023.0.2 \
    curl-minimal-8.5.0-1.amzn2023.0.5 \
    diffutils-3.8-1.amzn2023.0.2 \
    findutils-1:4.8.0-2.amzn2023.0.2 \
    git-2.47.1-1.amzn2023.0.3 \
    grep-3.8-1.amzn2023.0.4 \
    gzip-1.12-1.amzn2023.0.1 \
    iputils-20210202-2.amzn2023.0.4 \
    jq-1.7.1-49.amzn2023.0.2 \
    libxml2-2.10.4-1.amzn2023.0.10 \
    make-1:4.3-5.amzn2023.0.2 \
    openssh-clients-8.7p1-8.amzn2023.0.14 \
    patch-2.7.6-14.amzn2023.0.2 \
    pcre2-tools-10.40-1.amzn2023.0.3 \
    pwgen-2.08-11.amzn2023 \
    python3-3.9.22-1.amzn2023.0.1 \
    python3-pip-21.3.1-2.amzn2023.0.11 \
    rsync-3.4.0-1.amzn2023.0.2 \
    sed-4.8-7.amzn2023.0.2 \
    tar-2:1.34-1.amzn2023.0.4 \
    unzip-6.0-57.amzn2023.0.2 \
    wget-1.21.3-1.amzn2023.0.4 \
    xz-5.2.5-9.amzn2023.0.2 \
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
    json2yaml==1.2.0 \
    yamlpath==3.8.2 \
    yq==3.4.3 \
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
