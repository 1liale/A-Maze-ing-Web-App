# Default configuraitons provided by Hashicorp
terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "5.31.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

# EC2 Configuration
resource "aws_instance" "jenkins-ec2" {
  ami                    = "ami-079db87dc4c10ac91"
  instance_type          = "t2.micro"
  key_name               = "terraformkp"
  associate_public_ip_address = true
  vpc_security_group_ids = [aws_security_group.myjenkins_sg.id]
  user_data              = file("install.sh")
  tags = {
    Name = "Jenkins-server"
  }
}

# Create security group 
resource "aws_security_group" "myjenkins_sg" {
  name        = "jenkins-sg"
  description = "Allow Ports 22, 8080"

  # Allow Incoming SSH from Anywhere
  ingress {
    description = "Allow SSH Traffic"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Allow Incoming HTTPS from Anywhere
  ingress {
    description = "Allow HTTPS Traffic"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Allow Incoming HTTP from Anywhere
  ingress {
    description = "Allow HTTP Traffic"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Allow All Outgoing Requests
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Allocate S3 bucket to store Jenkins Artifacts
resource "aws_s3_bucket" "my-s3-bucket" {
  bucket = "jenkins-s3-bucket"

  tags = {
    Name = "Jenkins-server"
  }
}

# Define ACL for S3 Bucket
resource "aws_s3_bucket_acl" "s3_bucket_acl" {
  bucket     = aws_s3_bucket.my-s3-bucket.id
  acl        = "private"
  depends_on = [aws_s3_bucket_ownership_controls.s3_bucket_acl_ownership]
}

# Resource to avoid error "AccessControlListNotSupported: The bucket does not allow ACLs"
resource "aws_s3_bucket_ownership_controls" "s3_bucket_acl_ownership" {
  bucket = aws_s3_bucket.my-s3-bucket.id
  rule {
    object_ownership = "ObjectWriter"
  }
}