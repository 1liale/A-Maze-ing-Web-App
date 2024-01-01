data "aws_ami" "jenkins_master" {
  most_recent = true
  owners = ["self"]

  filter {
    name = "name"
    values = ["jenkins-master"]
  }
}

# Create security group 
resource "aws_security_group" "jenkins_master_sg" {
  name        = "jenkins_master_sg"
  description = "Allow Ports 22, 8080"
  vpc_id = "vpc-0036a7a64a87e6651"

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

  tags = {
    Name = "jenkins_master_sg"
  }
}

resource "aws_instance" "jenkins_master" {
  ami                    = data.aws_ami.jenkins_master.id
  instance_type          = var.instance_type
  key_name               = var.key
  vpc_security_group_ids = [aws_security_group.jenkins_master_sg.id]

  root_block_device {
    volume_type           = "gp3"
    volume_size           = 10
    delete_on_termination = false
  }

  tags = {
    Name   = "jenkins_master"
  }
}