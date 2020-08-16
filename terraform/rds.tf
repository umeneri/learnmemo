resource "aws_security_group" "praivate_db_sg" {
  name = "${var.env}-${var.task_name}-db-sg"
  vpc_id = aws_vpc.app_vpc.id

  ingress {
    from_port = 3306
    to_port = 3306
    protocol = "tcp"
    cidr_blocks = [
      aws_subnet.app_public_subnet_a.cidr_block,
      aws_subnet.app_public_subnet_c.cidr_block,
      aws_subnet.app_private_subnet_a.cidr_block,
      aws_subnet.app_private_subnet_c.cidr_block,
    ]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = [
      "0.0.0.0/0"]
  }
  tags = {
    Name = "${var.env}-${var.task_name}-db-sg"
  }
}
resource "aws_db_subnet_group" "praivate_db" {
  name = "${var.env}-${var.task_name}-db-group"
  subnet_ids = [
    aws_subnet.app_private_subnet_a.id,
    aws_subnet.app_private_subnet_c.id]
  tags = {
    Name = "${var.env}-${var.task_name}-db-group"
  }
}

resource "aws_db_instance" "db" {
  identifier = "${var.env}-${var.task_name}-db"
  allocated_storage = 20
  storage_type = "gp2"
  engine = "mysql"
  engine_version = "5.7"
  instance_class = "db.t3.micro"
  name = var.task_name
  username = var.task_name
  password = random_password.password.result
  parameter_group_name = "default.mysql5.7"
  vpc_security_group_ids = [
    aws_security_group.praivate_db_sg.id]
  db_subnet_group_name = aws_db_subnet_group.praivate_db.name
  skip_final_snapshot = true
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "@"
}