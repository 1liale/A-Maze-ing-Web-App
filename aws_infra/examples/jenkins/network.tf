# Define VPC to host jenkins cluster
resource "aws_vpc" "jenkins_vpc" {
  cidr_block = var.vpc_cidr_block

  tags = {
    Name = "jenkins-vpc"
  }
}

# Define public subnet
resource "aws_subnet" "public_subnet" {
  vpc_id                  = aws_vpc.jenkins_vpc.id
  cidr_block              = var.subnet_cidr_block
  availability_zone       = var.availability_zone

  tags = {
    Name = "jenkins-public-subnet"
  }
}

# Define internet gateway
resource "aws_internet_gateway" "jenkins_internet_gateway" {
  vpc_id = aws_vpc.jenkins_vpc.id

  tags = {
    Name = "jenkins-internet-gateway"
  }
}

# Define AWS route table
resource "aws_route_table" "jenkins_rtb" {
  vpc_id = aws_vpc.jenkins_vpc.id

  # Handle ipv4 routing
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.jenkins_internet_gateway.id
  }

  # Handle ipv6 routing
  route {
    ipv6_cidr_block = "::/0"
    gateway_id      = aws_internet_gateway.jenkins_internet_gateway.id
  }

  tags = {
    Name = "jenkins-rtb"
  }
}