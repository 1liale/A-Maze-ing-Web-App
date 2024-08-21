variable "region" {
  type        = string
  description = "AWS region"
}

variable "instance_type" {
  type        = string
  description = "EC2 instance type"
  default     = "t2.micro"
}

variable "jenkins_username" {
  type        = string
  description = "Jenkins admin user"
}

variable "jenkins_password" {
  type        = string
  description = "Jenkins admin password"
}

variable "jenkins_credentials_id" {
  type        = string
  description = "Jenkins workers SSH based credentials id"
  default     = "jenkins-workers"
}

variable "key" {
  type        = string
  description = "SSH key pair"
}

variable "vpc_cidr_block" {
  type        = string
  description = "Block of vpc addresses"
}

variable "subnet_cidr_block" {
  type        = string
  description = "Block of subnet addresses"
}

variable "availability_zone" {
  type        = string
  description = "Preferred availability zone"
}

variable "postgres_identifier" {
  type = string
}

variable "postgres_name" {
  type = string
}

variable "postgres_user_name" {
  type = string
}

variable "postgres_user_password" {
  type = string
}

variable "postgres_instance_name" {
  type = string
}

variable "postgres_db_password" {
  type = string
}

variable "postgres_port" {
  type = string
}
