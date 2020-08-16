resource "aws_ssm_parameter" "rds_password" {
  name   = "/rds/${var.task_name}/password"
  type   = "SecureString"
  key_id =  data.aws_kms_alias.ssm.target_key_id
  value  = aws_db_instance.db.password
}

data "aws_kms_alias" "ssm" {
  name = "alias/aws/ssm"
}