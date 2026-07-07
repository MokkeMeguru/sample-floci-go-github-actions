package usecase

import (
	"context"
	"fmt"

	"github.com/MokkeMeguru/sample-floci-go-github-actions/internal/pkg/asset"
	"github.com/aws/aws-sdk-go-v2/aws"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3API interface {
	CreateBucket(ctx context.Context, params *awss3.CreateBucketInput, optFns ...func(*awss3.Options)) (*awss3.CreateBucketOutput, error)
	DeleteBucket(ctx context.Context, params *awss3.DeleteBucketInput, optFns ...func(*awss3.Options)) (*awss3.DeleteBucketOutput, error)
}

type VerifyS3 struct {
	s3     S3API
	assets asset.Client
}

func NewVerifyS3(s3 S3API, assets asset.Client) *VerifyS3 {
	return &VerifyS3{s3: s3, assets: assets}
}

func (u *VerifyS3) Execute(ctx context.Context, bucket, key string) error {
	if _, err := u.s3.CreateBucket(ctx, &awss3.CreateBucketInput{
		Bucket: aws.String(bucket),
	}); err != nil {
		return fmt.Errorf("create bucket: %w", err)
	}
	defer u.s3.DeleteBucket(context.Background(), &awss3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})

	if err := u.assets.Put(ctx, bucket, key, "verified"); err != nil {
		return fmt.Errorf("put asset: %w", err)
	}
	if err := u.assets.Head(ctx, bucket, key); err != nil {
		return fmt.Errorf("head asset: %w", err)
	}
	if err := u.assets.Delete(ctx, bucket, key); err != nil {
		return fmt.Errorf("delete asset: %w", err)
	}
	return nil
}
