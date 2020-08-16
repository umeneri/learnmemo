resource "aws_ecr_repository" "app_repository" {
  name = "${var.env}-${var.task_name}"
}

