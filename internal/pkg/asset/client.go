package asset

import (
	"context"
	"strings"

	infras3 "github.com/MokkeMeguru/sample-floci-go-github-actions/internal/infrastructure/s3"
	"github.com/aws/aws-sdk-go-v2/aws"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client interface {
	Put(ctx context.Context, bucket, key, body string) error
	Head(ctx context.Context, bucket, key string) error
	Delete(ctx context.Context, bucket, key string) error
}

type client struct {
	s3 *infras3.Client
}

func NewClient(s3Client *infras3.Client) Client {
	return &client{s3: s3Client}
}

func (c *client) Put(ctx context.Context, bucket, key, body string) error {
	_, err := c.s3.S3.PutObject(ctx, &awss3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(body),
	})
	return err
}

func (c *client) Head(ctx context.Context, bucket, key string) error {
	_, err := c.s3.S3.HeadObject(ctx, &awss3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return err
}

func (c *client) Delete(ctx context.Context, bucket, key string) error {
	_, err := c.s3.S3.DeleteObject(ctx, &awss3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return err
}
