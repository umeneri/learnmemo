resource "aws_ecs_cluster" "app_ecs_cluster" {
  name = "${var.env}-${var.task_name}-cluster"
}

resource "aws_ecs_service" "app_ecs_app_service" {
  name = "${var.env}-${var.task_name}-service"
  cluster = aws_ecs_cluster.app_ecs_cluster.id
  task_definition = "${var.env}-${var.task_name}"
  desired_count = 1
  launch_type = "FARGATE"
  deployment_minimum_healthy_percent = 100
  deployment_maximum_percent = 200
  health_check_grace_period_seconds = 30
  depends_on = [
    aws_lb.app_lb
  ]

  network_configuration {
    subnets = [
      aws_subnet.app_private_subnet_a.id,
      aws_subnet.app_private_subnet_c.id
    ]
    security_groups = [
      aws_security_group.app_sg.id
    ]
    assign_public_ip = false
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.app_target_group.arn
    container_name = "learnmemo"
    container_port = 8080
  }

  lifecycle {
    ignore_changes = [
      task_definition
    ]
  }
}

resource "aws_cloudwatch_log_group" "app_log" {
  name = "/ecs/${var.env}-${var.task_name}"
  retention_in_days = 7
}
