variable "region" {
  type = string
  description = "AWS region"
}

variable "instance_type" {
  type = string
  description = "EC2 instance type"
  default = "t2.micro"
}

variable "jenkins_username" {
  type = string
  description = "Jenkins admin user"
}

variable "jenkins_password" {
  type = string
  description = "Jenkins admin password"
}

variable "jenkins_credentials_id" {
  type = string
  description = "Jenkins workers SSH based credentials id"
  default = "jenkins-workers"
}

variable "key" {
  type        = string
  description = "SSH key pair"
}

variable vpc_cidr_block {
  type = string
  description = "Block of vpc addresses"
}

variable subnet_cidr_block {
  type = string
  description = "Block of subnet addresses"
}          

variable "availability_zone" {
  type = string
  description = "Preferred availability zone"
}

variable "my_ip" {
  type = string
  description = "Limit jenkins SSH connection to my IP for better security"
}