# AWS CLI

Docker image to run the AWS CLI and related tools.

[![Dockerfile](https://img.shields.io/badge/GitHub-Dockerfile-blue)](https://github.com/leplusorg/docker-aws-cli/blob/main/aws-cli/Dockerfile)
[![Docker Build](https://github.com/leplusorg/docker-aws-cli/workflows/Docker/badge.svg)](https://github.com/leplusorg/docker-aws-cli/actions?query=workflow:"Docker")
[![Docker Stars](https://img.shields.io/docker/stars/leplusorg/aws-cli)](https://hub.docker.com/r/leplusorg/aws-cli)
[![Docker Pulls](https://img.shields.io/docker/pulls/leplusorg/aws-cli)](https://hub.docker.com/r/leplusorg/aws-cli)
[![Docker Version](https://img.shields.io/docker/v/leplusorg/aws-cli?sort=semver)](https://hub.docker.com/r/leplusorg/aws-cli)

## Usage

To run the AWS CLI using this image:

```bash
docker run --rm -i leplusorg/aws-cli aws --version
```

## Software Bill of Materials (SBOM)

To get the SBOM for the latest image (in SPDX JSON format), use the
following command:

```bash
docker buildx imagetools inspect leplusorg/aws-cli --format '{{ json (index .SBOM "linux/amd64").SPDX }}'
```

Replace `linux/amd64` by the desired platform (`linux/amd64`, `linux/arm64` etc.).
