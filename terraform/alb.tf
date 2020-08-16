resource "aws_lb" "app_lb" {
  name = "${var.env}-${var.task_name}-alb"
  internal = false
  load_balancer_type = "application"
  security_groups = [
    aws_security_group.alb_sg.id]
  subnets = [
    aws_subnet.app_public_subnet_a.id,
    aws_subnet.app_public_subnet_c.id,
  ]
  enable_deletion_protection = false
}

resource "aws_lb_listener" "app_listener" {
  load_balancer_arn = aws_lb.app_lb.arn
  certificate_arn = aws_acm_certificate.cert.arn
  port              = 443
  protocol          = "HTTPS"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.app_target_group.arn
  }
}

resource "aws_lb_target_group" "app_target_group" {
  name = "${var.env}-${var.task_name}-lb-tg"
  port = 8080
  protocol = "HTTP"
  vpc_id = aws_vpc.app_vpc.id
  target_type = "ip"

  health_check {
    interval            = 30
    path                = "/health"
    port                = 8080
    protocol            = "HTTP"
    timeout             = 5
    unhealthy_threshold = 2
    matcher             = 200
  }
}