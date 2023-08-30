FROM amazon/aws-cli:2.13.13

# hadolint ignore=DL3033
RUN yum update -y \
    && yum install -y bash curl git jq make python-pip3 tar unzip xmlstarlet zip \
    && yum clean all \
    && rm -rf /var/cache/yum

RUN pip3 install cfn-policy-validator

WORKDIR /opt

ENTRYPOINT []

CMD ["/bin/bash"]
