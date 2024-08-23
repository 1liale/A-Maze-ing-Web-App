# Define association to route internet traffic to public subnet
resource "aws_route_table_association" "public_rt_a" {
  subnet_id = aws_subnet.public_subnet.id
  route_table_id = aws_route_table.jenkins_rtb.id
}

# Define security group for the jenkins controller
resource "aws_security_group" "jenkins_master_sg" {
  name        = "jenkins-master-sg"
  description = "Allow Ports 22, 8080"
  vpc_id = aws_vpc.jenkins_vpc.id

  # Allow Incoming SSH from Anywhere
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Allow Incoming port 8080 from Anywhere
  ingress {
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

# Define jenkins-master AMI
data "aws_ami" "jenkins_master" {
  most_recent = true
  owners = ["self"]

  filter {
    name = "name"
    values = ["jenkins-master"]
  }
}

# Define jenkins-master ec2 instance
resource "aws_instance" "jenkins_master" {
  ami                    = data.aws_ami.jenkins_master.id
  instance_type          = var.instance_type
  key_name               = var.key

  subnet_id = aws_subnet.public_subnet.id
  vpc_security_group_ids = [aws_security_group.jenkins_master_sg.id]
  associate_public_ip_address = true

  root_block_device {
    volume_type           = "gp3"
    volume_size           = 30
    delete_on_termination = true
  }

  tags = {
    Name   = "jenkins-master"
  }
}