FROM amazon/aws-cli:2.24.27@sha256:d6a87ef206b1f6c1038dd583855c2af0c458246605e99a1c2493c98b04e1f0ba

HEALTHCHECK NONE

ARG USER_NAME=default
ARG USER_HOME=/home/default
ARG USER_ID=1000

SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

# Upgrade python to 3.8 (best we can do with Amazon Linux 2)
RUN yum remove python3 \
    && amazon-linux-extras install python3.8 \
    && ln -s /usr/bin/python3.8 /usr/bin/python3

RUN yum update -y \
    && yum install -y --setopt=skip_missing_names_on_install=False \
    bash-4.2.46-34.amzn2 \
    curl-8.3.0-1.amzn2.0.8 \
    git-2.47.1-1.amzn2.0.2 \
    jq-1.5-1.amzn2.0.2 \
    make-3.82-24.amzn2 \
    python3-pip-20.2.2-1.amzn2.0.8 \
    tar-1.26-35.amzn2.0.4 \
    unzip-6.0-57.amzn2.0.1 \
    zip-3.0-11.amzn2.0.2 \
    && IFS=$'\n\t' \
    && if yum list updates | grep -q -e '^Updated Packages' ; then \
       yum list updates ; \
       exit 1 ; \
    fi \
    && yum clean all \
    && rm -rf /var/cache/yum

# @TL FIXME: cfn-policy-validator version is held back by Python 3.8.
# @TL FIXME: --break-system-packages not supported until Python 3.11.
RUN pip3 install --no-cache-dir --upgrade \
    cfn-policy-validator==0.0.29 \
    && IFS=$'\n\t' \
    && msg="$(pip3 list --outdated | grep -i -e '^cfn-policy-validator ' || true)" \
    && if [ -n "${msg}" ]; then \
       >&2 echo "ERROR: outdated: ${msg}" ; \
       exit 1 ; \
    fi

RUN chmod 777 /opt \
    && adduser --home-dir "${USER_HOME}" --uid "${USER_ID}" "${USER_NAME}"

ENV HOME="${USER_HOME}"

USER "${USER_NAME}"

WORKDIR /opt

CMD ["/bin/bash"]

ENTRYPOINT []
