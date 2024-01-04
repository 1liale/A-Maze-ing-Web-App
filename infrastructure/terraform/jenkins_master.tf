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
  vpc_id = aws_vpc.jenkins_vpc.id

  # Allow Incoming SSH from Anywhere
  ingress {
    description = "Allow SSH Traffic"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [var.my_ip]
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
    Name = "jenkins-master-sg"
  }
}

resource "aws_instance" "jenkins_master" {
  ami                    = data.aws_ami.jenkins_master.id
  instance_type          = var.instance_type
  key_name               = var.key

  vpc_security_group_ids = [aws_security_group.jenkins_master_sg.id]

  subnet_id = aws_subnet.public_subnet.id
  associate_public_ip_address = true

  root_block_device {
    volume_type           = "gp3"
    volume_size           = 15
    delete_on_termination = true
  }

  tags = {
    Name   = "jenkins-master"
  }
}