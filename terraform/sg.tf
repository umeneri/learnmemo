resource "aws_security_group" "alb_sg" {
  name = "${var.env}-${var.task_name}-alb-sg"
  vpc_id = aws_vpc.app_vpc.id

  tags = {
    Name = "${var.env}-${var.task_name}-sg"
  }
}

resource "aws_security_group_rule" "alb_ingress" {
  type = "ingress"
  from_port = 443
  to_port = 443
  protocol = "tcp"
  cidr_blocks = [
    "0.0.0.0/0"]

  security_group_id = aws_security_group.alb_sg.id
}

resource "aws_security_group_rule" "alb_egress" {
  type = "egress"
  from_port = 0
  to_port = 0
  protocol = "-1"
  cidr_blocks = [
    "0.0.0.0/0"]

  security_group_id = aws_security_group.alb_sg.id
}

resource "aws_security_group" "app_sg" {
  name = "${var.env}-${var.task_name}-sg"
  vpc_id = aws_vpc.app_vpc.id

  tags = {
    Name = "${var.env}-${var.task_name}-sg"
  }
}

resource "aws_security_group_rule" "app_ingress" {
  type = "ingress"
  from_port = 8080
  to_port = 8080
  protocol = "tcp"
  cidr_blocks = [
    "0.0.0.0/0"]

  security_group_id = aws_security_group.app_sg.id
}

resource "aws_security_group_rule" "app_egress" {
  type = "egress"
  from_port = 0
  to_port = 0
  protocol = "-1"
  cidr_blocks = [
    "0.0.0.0/0"]

  security_group_id = aws_security_group.app_sg.id
}
