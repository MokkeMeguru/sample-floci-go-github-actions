# sample-floci-go-github-actions

Minimal Go sample for verifying Floci S3 on GitHub Actions without copying private application code.

The sample keeps a small layered structure:

- `internal/config`: environment-based settings
- `internal/infrastructure/s3`: AWS SDK S3 client construction
- `internal/pkg/asset`: S3 object operations
- `internal/usecase`: bucket/object verification flow

GitHub Actions uses `AWS_ENDPOINT_URL_S3=http://localhost.floci.io:4566`.
With Go SDK v2 default S3 addressing, requests become virtual-hosted-style URLs such as:

```text
http://<bucket>.localhost.floci.io:4566/<key>
```

This keeps local emulator verification closer to production S3 than forcing path-style access.

## Run locally

```sh
docker run --rm -p 4566:4566 \
  -e AWS_ACCESS_KEY_ID=test \
  -e AWS_SECRET_ACCESS_KEY=local-test-secret \
  -e AWS_DEFAULT_REGION=ap-northeast-1 \
  -e FLOCI_DEFAULT_REGION=ap-northeast-1 \
  floci/floci:1.5.30-compat
```

In another shell:

```sh
FLOCI_INTEGRATION_TEST=1 \
AWS_ACCESS_KEY_ID=test \
AWS_SECRET_ACCESS_KEY=local-test-secret \
AWS_DEFAULT_REGION=ap-northeast-1 \
AWS_ENDPOINT_URL_S3=http://localhost.floci.io:4566 \
go test ./...
```
