# Prepare Go Lambda artifact using the module's built-in packaging pipeline.
locals {
  project_root = "${path.module}/.."
  artifact_dir = "${path.module}/build"
}

module "lambda_api" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "8.1.2"

  function_name = var.function_name
  handler       = var.handler
  runtime       = var.runtime
  architectures = var.architectures

  create_package = true
  source_path = [
    {
      path = local.project_root
      commands = [
        "set -e",
        "rm -rf \"${local.artifact_dir}\"",
        "mkdir -p \"${local.artifact_dir}\"",
        "GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o \"${local.artifact_dir}/bootstrap\" main.go",
        ":zip ${local.artifact_dir}",
      ]
    }
  ]

  environment_variables = {
    LINE_CHANNEL_ACCESS_TOKEN = var.line_channel_access_token
    LINE_TARGET_ID            = var.line_target_id
  }

  tags = try(var.tags, {})
}
