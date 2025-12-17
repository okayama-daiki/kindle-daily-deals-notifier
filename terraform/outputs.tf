output "lambda_function_name" {
  value       = module.lambda_api.lambda_function_name
  description = "Name of the deployed Lambda function"
}

output "lambda_function_arn" {
  value       = module.lambda_api.lambda_function_arn
  description = "ARN of the deployed Lambda function"
}

output "scheduler_schedule_arn" {
  value       = try(aws_scheduler_schedule.daily_invoke_lambda[0].arn, null)
  description = "ARN of the EventBridge Scheduler schedule (if created)"
}
