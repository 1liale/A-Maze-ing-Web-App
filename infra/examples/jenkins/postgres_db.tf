resource "aws_security_group" "postgres_sg" {
  name = "postgres_sg"

  // default communication from port :5432
  ingress {
    from_port   = var.postgres_port
    to_port     = var.postgres_port
    protocol    = "tcp"
    description = "PostgreSQL"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port        = var.postgres_port
    to_port          = var.postgres_port
    protocol         = "tcp"
    description      = "PostgreSQL"
    ipv6_cidr_blocks = ["::/0"]
  }
}

resource "aws_db_instance" "postgres_rds" {
  db_name                = var.postgres_instance_name
  allocated_storage      = 20
  storage_type           = "gp2"
  engine                 = "postgres"
  engine_version         = "12.2"
  instance_class         = "db.t2.micro"
  identifier             = var.postgres_identifier
  username               = var.postgres_user_name
  password               = var.postgres_db_password
  publicly_accessible    = true
  parameter_group_name   = "default.postgres12"
  vpc_security_group_ids = [aws_security_group.postgres_sg.id]
  skip_final_snapshot    = true
}

resource "postgresql_role" "db_admin_role" {
  name                = "db_admin_role"
  login               = true
  password            = var.postgres_user_password
  encrypted_password  = true
  create_database     = true
  create_role         = true
  skip_reassign_owned = true
}
