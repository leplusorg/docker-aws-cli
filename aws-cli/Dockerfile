FROM amazon/aws-cli:2.4.13

ENV NODE_ENV production

RUN yum update -y \
    && yum install -y bash curl git jq make tar unzip zip \
    && yum clean all \
    && rm -rf /var/cache/yum \
    && curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash \
    && source ~/.nvm/nvm.sh \
    && nvm install --lts \
    && nvm use --lts

WORKDIR /opt

ENTRYPOINT []

CMD '/bin/bash'
