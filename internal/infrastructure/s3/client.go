package s3

import (
	"context"

	"github.com/MokkeMeguru/sample-floci-go-github-actions/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	S3 *awss3.Client
}

func NewClient(ctx context.Context, conf config.Config) (*Client, error) {
	cfg, err := awsconfig.LoadDefaultConfig(
		ctx,
		awsconfig.WithRegion(conf.AWS.Region),
	)
	if err != nil {
		return nil, err
	}
	if conf.AWS.EndpointURLS3 != "" {
		cfg.BaseEndpoint = aws.String(conf.AWS.EndpointURLS3)
	}
	return NewClientFromConfig(cfg), nil
}

func NewClientFromConfig(cfg aws.Config) *Client {
	return &Client{S3: awss3.NewFromConfig(cfg)}
}
