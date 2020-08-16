resource "aws_iam_role" "app_ecs_task_execution_role" {
  name = "ecs-task-execution"
  assume_role_policy = data.aws_iam_policy_document.app_task_role_assume_role_policy_document.json
}

resource "aws_iam_role_policy_attachment" "app_ecs_task_execution_role_policy_attachment" {
  role = aws_iam_role.app_ecs_task_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role_policy" "app_task_execution_policy" {
  name   = "${var.env}-${var.task_name}-execution"
  role   = aws_iam_role.app_ecs_task_execution_role.id
  policy = data.aws_iam_policy_document.app_task_execution_policy_document.json
}

data "aws_iam_policy_document" "app_task_role_assume_role_policy_document" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "app_task_execution_policy_document" {
  statement {
    actions = [
      "ssm:GetParameters",
      "ssm:GetParameter"
    ]

    resources = [
      aws_ssm_parameter.rds_password.arn,
    ]
  }
}
