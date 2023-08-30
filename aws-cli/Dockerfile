FROM amazon/aws-cli:2.13.13

# hadolint ignore=DL3033
RUN yum update -y \
    && yum install -y bash curl git jq make python3-pip tar unzip xmlstarlet zip \
    && yum clean all \
    && rm -rf /var/cache/yum

RUN pip3 install -no-cache-dir cfn-policy-validator==0.0.21

WORKDIR /opt

ENTRYPOINT []

CMD ["/bin/bash"]
