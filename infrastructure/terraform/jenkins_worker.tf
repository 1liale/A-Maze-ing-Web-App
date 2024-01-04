data "aws_ami" "jenkins_worker" {
  most_recent = true
  owners = ["self"]

  filter {
    name = "name"
    values = ["jenkins-worker"]
  }
}

# Create security group 
resource "aws_security_group" "jenkins_worker_sg" {
  name        = "jenkins_worker_sg"
  description = "Allows only SSH connections from Jenkins master"
  vpc_id = aws_vpc.jenkins_vpc.id

  # Allow Incoming SSH from Jenkins master SG
  ingress {
    description = "Allow SSH Traffic"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    security_groups = [aws_security_group.jenkins_master_sg.id]
  }

  # Allow All Outgoing Requests
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "jenkins-worker-sg"
  }
}

resource "aws_iam_role" "jenkins_role" {
  name = "jenkins_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = "RoleForEC2"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_iam_instance_profile" "jenkins_instance_profile" {
  name = "jenkins-instance-profile"
  role = aws_iam_role.jenkins_role.name
}

# Specify worker launch configuration
resource "aws_launch_configuration" "jenkins_workers_launch_conf" {
  name            = "jenkins_workers_asg_conf"
  image_id        = data.aws_ami.jenkins_worker.id
  instance_type   = var.instance_type

  key_name        = var.key
  associate_public_ip_address = true

  iam_instance_profile = aws_iam_instance_profile.jenkins_instance_profile.name
  security_groups = [aws_security_group.jenkins_worker_sg.id]
  user_data       = templatefile("scripts/join-cluster.tftpl", {
      jenkins_url            = "http://${aws_instance.jenkins_master.private_ip}:8080"
      jenkins_username       = var.jenkins_username
      jenkins_password       = var.jenkins_password
      jenkins_credentials_id = var.jenkins_credentials_id
  }) 

  root_block_device {
    volume_type           = "gp3"
    volume_size           = 15
    delete_on_termination = true
  }

  lifecycle {
    create_before_destroy = true
  }
}