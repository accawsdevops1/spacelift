provider "aws" {
  region = "us-east-1"
}

# Create IAM Groups
resource "aws_iam_group" "group1" {
  name = "group1"
}

resource "aws_iam_group" "group2" {
  name = "group2"
}

# Create IAM Users
resource "aws_iam_user" "user1" {
  name = "user1"
}

resource "aws_iam_user" "user2" {
  name = "user2"
}

# Add users to groups
resource "aws_iam_user_group_membership" "user1_membership" {
  user = aws_iam_user.user1.name
  groups = [
    aws_iam_group.group1.name
  ]
}

resource "aws_iam_user_group_membership" "user2_membership" {
  user = aws_iam_user.user2.name
  groups = [
    aws_iam_group.group2.name
  ]
}

# Outputs for Terratest to verify
output "iam_user1_name" {
  value = aws_iam_user.user1.name
}

output "iam_user2_name" {
  value = aws_iam_user.user2.name
}

output "iam_group1_name" {
  value = aws_iam_group.group1.name
}

output "iam_group2_name" {
  value = aws_iam_group.group2.name
}
