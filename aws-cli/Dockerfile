FROM amazon/aws-cli:2.3.4

RUN yum update -y \
    && yum install -y bash \
    && yum clean all \
    && rm -rf /var/cache/yum

WORKDIR /opt

ENTRYPOINT []

CMD '/bin/bash'
