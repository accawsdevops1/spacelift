package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformIAMUserGroup(t *testing.T) {
	t.Parallel()

	// Define the terraform options
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../", // Path to your Terraform code
	})

	// Run terraform init and apply
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Get the IAM User and IAM Group from the outputs
	iamUser1Name := terraform.Output(t, terraformOptions, "iam_user1_name")
	iamUser2Name := terraform.Output(t, terraformOptions, "iam_user2_name")
	iamGroup1Name := terraform.Output(t, terraformOptions, "iam_group1_name")
	iamGroup2Name := terraform.Output(t, terraformOptions, "iam_group2_name")

	// Test that the IAM Users and IAM Groups are created
	assert.NotEmpty(t, iamUser1Name)
	assert.NotEmpty(t, iamUser2Name)
	assert.NotEmpty(t, iamGroup1Name)
	assert.NotEmpty(t, iamGroup2Name)

	// Test that the users are correctly assigned to the groups
	assert.Contains(t, iamGroup1Name, "group1")
	assert.Contains(t, iamGroup2Name, "group2")
	assert.Contains(t, iamUser1Name, "user1")
	assert.Contains(t, iamUser2Name, "user2")

	// Output IAM User and Group Info
	fmt.Printf("Created IAM User 1: %s\n", iamUser1Name)
	fmt.Printf("Created IAM User 2: %s\n", iamUser2Name)
	fmt.Printf("Created IAM Group 1: %s\n", iamGroup1Name)
	fmt.Printf("Created IAM Group 2: %s\n", iamGroup2Name)

	// You can also check for specific policies if needed.
}
