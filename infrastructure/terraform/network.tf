# Define VPC to host jenkins cluster
resource "aws_vpc" "jenkins_vpc" {
  cidr_block = var.vpc_cidr_block
  enable_dns_support = true
  enable_dns_hostnames = true

  tags = {
    Name = "jenkins-vpc"
  }
}

# Define public subnet
resource "aws_subnet" "public_subnet" {
  vpc_id                  = aws_vpc.jenkins_vpc.id
  cidr_block              = var.subnet_cidr_block
  availability_zone       = var.availability_zone
  map_public_ip_on_launch = true

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

resource "aws_default_route_table" "jenkins_default_rtb" {
  default_route_table_id = aws_vpc.jenkins_vpc.default_route_table_id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.jenkins_internet_gateway.id
  }

  tags = {
    Name = "jenkins-default-rtb"
  }
}