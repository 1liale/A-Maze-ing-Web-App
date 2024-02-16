# Default configuraitons provided by Hashicorp
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.31.0"
    }
  }
}

provider "aws" {
  region = var.region
}

provider "postgresql" {
  host            = aws_db_instance.postgres.address
  port            = var.postgres_port
  database        = var.postgres_name
  username        = var.postgres_username
  password        = var.postgres_password
  sslmode         = "require"
  connect_timeout = 15
  superuser       = false
}
