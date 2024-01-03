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
    Name = "jenkins_worker_sg"
  }
}

# Specify worker launch configuration
resource "aws_launch_configuration" "jenkins_workers_launch_conf" {
  name            = "jenkins_workers_asg_conf"
  image_id        = data.aws_ami.jenkins_worker.id
  instance_type   = var.instance_type
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
    volume_size           = 10
    delete_on_termination = false
  }

  lifecycle {
    create_before_destroy = true
  }
}

# Auto-scaling
resource "aws_autoscaling_group" "jenkins_workers" {
  name                 = "jenkins_workers_asg"
  launch_configuration = aws_launch_configuration.jenkins_workers_launch_conf.name
  availability_zones = ["us-east-1d"]
  min_size             = 1
  desired_capacity     = 1
  max_size             = 4

  depends_on = [aws_instance.jenkins_master]

  lifecycle {
    create_before_destroy = true
  }

  tag {
    key                 = "Name"
    value               = "jenkins_worker"
    propagate_at_launch = true
  }
}