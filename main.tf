terraform {
  required_providers {
    aws = {
        source = "hashicorp/aws"
        version = "~> 4.21.0"
    }
    archive = {
        source = "hashicorp/archive"
        version = "~> 2.2.0"
    }
  }
  required_version = "~> 1.0"
}

provider "aws" {
  region = "ap-northeast-1"
  shared_config_files = [ "$HOME/.aws/credentials" ]
}


resource "aws_s3_bucket" "lambda_bucket" {

    bucket_prefix = "s3-chatops-lambda"
    tags = {
      name = "chatops-lambda-bucket"
      Environment = "Dev"
    }
    force_destroy = true
}

resource "aws_s3_bucket_public_access_block" "lambda_s3_role" {
  bucket = aws_s3_bucket.lambda_bucket.id
  block_public_acls = true
  block_public_policy = true
  restrict_public_buckets = true
  ignore_public_acls = true
}

resource "aws_iam_role" "lambda_execution" {

    name = "chatops_lambda_execution_test_role"
    tags = {tag-key = "test-lambda"}  

    assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      }, 
      "Action": "sts:AssumeRole"
    }
  ]
}
POLICY
    
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role = aws_iam_role.lambda_execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"

}


data "archive_file" "chatops_lambda_functions" {
  type = "zip"

  source_dir = "${path.module}/code"
  output_path = "${path.module}/code.zip"

}

resource "aws_cloudformation_stack" "chatops_lambda_stack" {
    name = "chatops-lambda-stack"

    template_body = file("${path.module}/template.yaml")
    capabilities   = ["CAPABILITY_AUTO_EXPAND"]
    

    tags = {
        Name = "chatops_lambda_stack"
    }
  
}

resource "aws_s3_object" "chatops_mern" {
  bucket =  aws_s3_bucket.lambda_bucket.id

  key = "code.zip"
  source = data.archive_file.chatops_lambda_functions.output_path

  etag = filemd5(data.archive_file.chatops_lambda_functions.output_path)

}

resource "aws_lambda_function" "lambda_functions" {
  function_name = "chatops_application"
  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key = aws_s3_object.chatops_mern.id

  role = aws_iam_role.lambda_execution.arn

  runtime = "go1.x" 
  handler = "main"

  source_code_hash = data.archive_file.chatops_lambda_functions.output_base64sha256

}
