package s3_test

import (
	"context"
	"net/http"
	"testing"

	infras3 "github.com/MokkeMeguru/sample-floci-go-github-actions/internal/infrastructure/s3"
	"github.com/aws/aws-sdk-go-v2/aws"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

type captureClient struct {
	request *http.Request
}

func (c *captureClient) Do(req *http.Request) (*http.Response, error) {
	c.request = req
	return &http.Response{
		StatusCode: http.StatusNoContent,
		Body:       http.NoBody,
		Header:     make(http.Header),
		Request:    req,
	}, nil
}

func TestNewClientFromConfigUsesVirtualHostedStyleWithEndpointOverride(t *testing.T) {
	httpClient := &captureClient{}
	client := infras3.NewClientFromConfig(aws.Config{
		Region:       "ap-northeast-1",
		BaseEndpoint: aws.String("http://s3.localhost.floci.io:4566"),
		HTTPClient:   httpClient,
	})

	_, err := client.S3.DeleteObject(context.Background(), &awss3.DeleteObjectInput{
		Bucket: aws.String("sample-bucket"),
		Key:    aws.String("path/to/object.txt"),
	})
	if err != nil {
		t.Fatalf("DeleteObject() error = %v", err)
	}
	if httpClient.request == nil {
		t.Fatal("expected request to be captured")
	}

	got := httpClient.request.URL.String()
	want := "http://sample-bucket.s3.localhost.floci.io:4566/path/to/object.txt?x-id=DeleteObject"
	if got != want {
		t.Fatalf("request URL = %q, want %q", got, want)
	}
}
