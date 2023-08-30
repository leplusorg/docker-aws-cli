FROM amazon/aws-cli:2.13.13

# hadolint ignore=DL3033
RUN yum update -y \
    && yum install -y bash curl git jq make openssl11 openssl11-devel python3-pip tar unzip xmlstarlet zip \
    && yum -y purge openssl-devel \
    && yum clean all \
    && rm -rf /var/cache/yum

# hadolint ignore=DL3013
RUN pip3 install --no-cache-dir cfn-policy-validator

WORKDIR /opt

ENTRYPOINT []

CMD ["/bin/bash"]
