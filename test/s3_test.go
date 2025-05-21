package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3Bucket(t *testing.T) {
	t.Parallel()

	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../", // Make sure this path is correct
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	bucketName := terraform.Output(t, terraformOptions, "bucket_name")

	// ✅ Correct method to assert the bucket exists
	aws.AssertS3BucketExists(t, awsRegion, bucketName)

	// Optional: Still asserting manually for demonstration
	assert.NotEmpty(t, bucketName, "Bucket name should not be empty")
}
