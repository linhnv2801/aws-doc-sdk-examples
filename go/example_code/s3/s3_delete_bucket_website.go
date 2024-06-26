// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Deletes the bucket's website configuration. Allows setting the index suffix,
// and an optional error page keys. If the bucket does not have a website
// configuration no error will be returned.
//
// If the bucket already has a website configured on it this will overwrite
// that configuration
//
// Usage:
//
//	go run s3_delete_bucket_website.go BUCKET_NAME
func main() {
	if len(os.Args) != 2 {
		exitErrorf("bucket name required\nUsage: %s bucket_name",
			filepath.Base(os.Args[0]))
	}

	bucket := os.Args[1]

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	// Deletes the website configuration on the bucket. Will return successfully
	// when the website configuration was deleted, or if the bucket does not
	// have a website configuration.
	_, err = svc.DeleteBucketWebsite(&s3.DeleteBucketWebsiteInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Unable to delete bucket %q website configuration, %v",
			bucket, err)
	}

	fmt.Printf("Successfully delete bucket %q website configuration\n", bucket)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
