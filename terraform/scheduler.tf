resource "aws_iam_role" "eventbridge_scheduler_invoke_lambda" {
  count = var.create_scheduler ? 1 : 0
  name  = "${var.function_name}-scheduler-invoke-lambda"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "scheduler.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })

  tags = try(var.tags, {})
}

resource "aws_iam_role_policy" "eventbridge_scheduler_invoke_lambda" {
  count = var.create_scheduler ? 1 : 0
  name  = "${var.function_name}-scheduler-invoke-lambda"
  role  = aws_iam_role.eventbridge_scheduler_invoke_lambda[0].id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = ["lambda:InvokeFunction"]
        Resource = [module.lambda_api.lambda_function_arn]
      }
    ]
  })
}

resource "aws_scheduler_schedule" "daily_invoke_lambda" {
  count = var.create_scheduler ? 1 : 0

  name                         = var.scheduler_name
  schedule_expression          = var.scheduler_schedule_expression
  schedule_expression_timezone = var.scheduler_schedule_timezone

  flexible_time_window {
    mode = "OFF"
  }

  target {
    arn      = module.lambda_api.lambda_function_arn
    role_arn = aws_iam_role.eventbridge_scheduler_invoke_lambda[0].arn
  }
}
