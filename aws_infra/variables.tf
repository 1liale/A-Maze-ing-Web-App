variable "region" {
  type        = string
  description = "AWS region"
  default     = "us-east-1"
}

variable "app_name" {
  description = "Application name"
  default     = "A-Maze-ing-lambda-api"
}

variable "app_env" {
  description = "Application environment tag"
  default     = "dev"
}
