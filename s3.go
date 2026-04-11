package main

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	*s3.Client
}

func NewS3Client(cfg aws.Config) *S3Client {
	return &S3Client{s3.NewFromConfig(cfg)}
}

func (c *S3Client) downloadFromS3(ctx context.Context, bucket, key, dest string) error {
	object, err := c.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	defer object.Body.Close()

	file, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, object.Body)
	return err
}

func (c *S3Client) uploadToS3(ctx context.Context, bucket, key, src string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = c.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	return err
}
