variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "ap-northeast-1"
}

variable "function_name" {
  description = "Name of the Lambda function."
  type        = string
  default     = "notifyKindleDailyDeals"
}

variable "line_channel_access_token" {
  description = "LINE channel access token used by the Lambda (sensitive). Prefer provisioning via CI secrets or a secret manager."
  type        = string
  sensitive   = true
  default     = ""
}

variable "line_target_id" {
  description = "LINE target ID (user/group/room) to push messages to (sensitive)."
  type        = string
  sensitive   = true
  default     = ""
}

variable "handler" {
  description = "Lambda handler. For custom runtimes using a bootstrap binary, use 'bootstrap'."
  type        = string
  default     = "bootstrap"
}

variable "runtime" {
  description = "Lambda runtime. Use 'provided.al2023' for custom Amazon Linux 2023 runtime."
  type        = string
  default     = "provided.al2023"
}

variable "architectures" {
  description = "CPU architecture(s) for the Lambda function."
  type        = list(string)
  default     = ["arm64"]
}

variable "tags" {
  description = "Map of tags to apply to created resources."
  type        = map(string)
  default     = {}
}

variable "create_scheduler" {
  description = "Whether to create an EventBridge Scheduler schedule to invoke the Lambda."
  type        = bool
  default     = true
}

variable "scheduler_name" {
  description = "Name of the EventBridge Scheduler schedule."
  type        = string
  default     = "notify-kindle-daily-deals-daily-7am-jst"
}

variable "scheduler_schedule_expression" {
  description = "Schedule expression for EventBridge Scheduler. Example: cron(0 7 * * ? *)."
  type        = string
  default     = "cron(0 7 * * ? *)"
}

variable "scheduler_schedule_timezone" {
  description = "Timezone for the schedule expression (e.g., Asia/Tokyo)."
  type        = string
  default     = "Asia/Tokyo"
}
