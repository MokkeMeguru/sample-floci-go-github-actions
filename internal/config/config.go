package config

import "os"

type Config struct {
	AWS AWSConfig
	S3  S3Config
}

type AWSConfig struct {
	EndpointURLS3 string
	Region        string
}

type S3Config struct {
	BucketPrefix string
}

func Load() Config {
	return Config{
		AWS: AWSConfig{
			EndpointURLS3: getenv("AWS_ENDPOINT_URL_S3", "http://s3.localhost.floci.io:4566"),
			Region:        getenv("AWS_DEFAULT_REGION", "ap-northeast-1"),
		},
		S3: S3Config{
			BucketPrefix: getenv("S3_BUCKET_PREFIX", "sample-floci"),
		},
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
