# aws-cli image

Build context for the `leplusorg/aws-cli` container image. See the
[root readme](../README.md) for what the image does and how to run it.

## Files

- `Dockerfile` — the image definition. System packages (including the AWS CLI
  itself) come from Amazon Linux's `dnf` and are pinned to exact versions here.
- `docker-compose.test.yml` — the smoke test run in CI against the built image:
  the `sut` service checks that the container runs as the non-root user (UID 1000) and that each bundled tool starts.
- `requirements.txt` — the Python tools installed in the image, with their
  pinned versions. The `Dockerfile` installs each one globally with `pipx`
  (using Python 3.14), in its own virtual environment.
- `go.mod` — the Go tools installed in the image, with their pinned versions.
  The `Dockerfile` runs `go install` for each direct `require` line (currently
  only `yq`, a workaround until Amazon Linux ships a `yq` package). The `//
indirect` entries are their transitive dependencies, maintained by `go mod
tidy`.
- `go.sum` — the checksums of the modules in `go.mod`, maintained by `go mod
tidy`. Not copied into the image.
- `main.go` — a dummy program that imports `yq` so that `go mod tidy` keeps it
  in `go.mod`. Not built into the image.

## Manifests

The manifests (`requirements.txt`, `go.mod`) expose the pinned versions of some
of the packages and tools included in the image in a machine-readable format.
They can be used to generate an SBOM, and they are kept up to date automatically
by Dependabot (see [`dependabot.yml`](../.github/dependabot.yml)); Renovate
would work too.
