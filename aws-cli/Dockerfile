FROM amazon/aws-cli

RUN yum update -y \
    && yum install -y bash \
    && yum clean all \
    && rm -rf /var/cache/yum

CMD ['bash']
