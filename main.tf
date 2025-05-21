provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "example" {
  bucket = "terratest-example-bucket-9876"  # Make it globally unique!
  force_destroy = true
}

output "bucket_name" {
  value = aws_s3_bucket.example.bucket
}
