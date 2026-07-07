package usecase_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/MokkeMeguru/sample-floci-go-github-actions/internal/config"
	infras3 "github.com/MokkeMeguru/sample-floci-go-github-actions/internal/infrastructure/s3"
	"github.com/MokkeMeguru/sample-floci-go-github-actions/internal/pkg/asset"
	"github.com/MokkeMeguru/sample-floci-go-github-actions/internal/usecase"
)

func TestVerifyS3WithFloci(t *testing.T) {
	if os.Getenv("FLOCI_INTEGRATION_TEST") != "1" {
		t.Skip("set FLOCI_INTEGRATION_TEST=1 to run Floci integration test")
	}

	ctx := context.Background()
	conf := config.Load()
	s3Client, err := infras3.NewClient(ctx, conf)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	bucket := fmt.Sprintf("%s-%d", conf.S3.BucketPrefix, time.Now().UnixNano())
	key := "push/thumbnail/addressing-test.txt"

	verify := usecase.NewVerifyS3(s3Client.S3, asset.NewClient(s3Client))
	if err := retry(ctx, 60, time.Second, func() error {
		return verify.Execute(ctx, bucket, key)
	}); err != nil {
		t.Fatalf("VerifyS3.Execute() error = %v", err)
	}
}

func retry(ctx context.Context, attempts int, wait time.Duration, fn func() error) error {
	var lastErr error
	for range attempts {
		if err := fn(); err == nil {
			return nil
		} else {
			lastErr = err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(wait):
		}
	}
	return lastErr
}
