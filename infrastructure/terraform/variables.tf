variable "region" {
  type = string
  description = "AWS region"
  default = "us-east-1"
}

variable "instance_type" {
  type = string
  description = "EC2 instance type"
  default = "t2.micro"
}

variable "jenkins_username" {
  type = string
  description = "Jenkins admin user"
  default = "admin"
}

variable "jenkins_password" {
  type = string
  description = "Jenkins admin password"
  default = "admin"
}

variable "jenkins_credentials_id" {
  type = string
  description = "Jenkins workers SSH based credentials id"
  default = "jenkins-workers"
}

variable "key" {
  type        = string
  description = "SSH key pair"
  default = "terraformkp"
}