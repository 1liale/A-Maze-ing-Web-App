# Define security group 
resource "aws_security_group" "jenkins_worker_sg" {
  name        = "jenkins-worker-sg"
  description = "Allows SSH from master"
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

data "aws_ami" "jenkins_worker" {
  most_recent = true
  owners = ["self"]

  filter {
    name = "name"
    values = ["jenkins-worker"]
  }
}

# Specify worker launch configuration
resource "aws_launch_configuration" "jenkins_workers_launch_conf" {
  name            = "jenkins-workers-asg-conf"
  image_id        = data.aws_ami.jenkins_worker.id
  instance_type   = var.instance_type
  key_name        = var.key
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
    volume_size           = 30
    delete_on_termination = true
  }

  lifecycle {
    create_before_destroy = true
  }
}