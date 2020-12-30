FROM amazon/aws-cli:latest

RUN yum update -y \
    && yum install -y bash \
    && yum clean all \
    && rm -rf /var/cache/yum

ENTRYPOINT ['/bin/bash']
