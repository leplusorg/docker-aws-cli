# AWS CLI

Docker image to run the AWS CLI and related tools.

[![Dockerfile](https://img.shields.io/badge/GitHub-Dockerfile-blue)](https://github.com/leplusorg/docker-aws-cli/blob/main/aws-cli/Dockerfile)
[![Docker Build](https://github.com/leplusorg/docker-aws-cli/workflows/Docker/badge.svg)](https://github.com/leplusorg/docker-aws-cli/actions?query=workflow:"Docker")
[![Docker Stars](https://img.shields.io/docker/stars/leplusorg/aws-cli)](https://hub.docker.com/r/leplusorg/aws-cli)
[![Docker Pulls](https://img.shields.io/docker/pulls/leplusorg/aws-cli)](https://hub.docker.com/r/leplusorg/aws-cli)
[![Docker Version](https://img.shields.io/docker/v/leplusorg/aws-cli?sort=semver)](https://hub.docker.com/r/leplusorg/aws-cli)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/10078/badge)](https://bestpractices.coreinfrastructure.org/projects/10078)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/leplusorg/docker-aws-cli/badge)](https://securityscorecards.dev/viewer/?uri=github.com/leplusorg/docker-aws-cli)

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

### Sigstore

[Sigstore](https://docs.sigstore.dev) is trying to improve supply
chain security by allowing you to verify the origin of an
artifcat. You can verify that the jar that you use was actually
produced by this repository. This means that if you verify the
signature of the ristretto jar, you can trust the integrity of the
whole supply chain from code source, to CI/CD build, to distribution
on Maven Central or whever you got the jar from.

You can use the following command to verify the latest image using its
sigstore signature attestation:

```bash
cosign verify leplusorg/aws-cli --certificate-identity-regexp 'https://github\.com/leplusorg/docker-aws-cli/\.github/workflows/.+' --certificate-oidc-issuer 'https://token.actions.githubusercontent.com'
```

The output should look something like this:

```text
Verification for index.docker.io/leplusorg/xml:main --
The following checks were performed on each of these signatures:
  - The cosign claims were validated
  - Existence of the claims in the transparency log was verified offline
  - The code-signing certificate was verified using trusted certificate authority certificates

[{"critical":...
```

For instructions on how to install `cosign`, please read this [documentation](https://docs.sigstore.dev/cosign/system_config/installation/).
