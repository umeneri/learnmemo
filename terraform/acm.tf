# [TerraformでACM証明書を作成してみた | Developers.IO](https://dev.classmethod.jp/articles/acm-cert-by-terraform/)
# [TerraformでACMの証明書を取得する方法 - Qiita](https://qiita.com/dd511805/items/b7546a4166c1bf290fc4)

resource aws_acm_certificate cert {
  domain_name       = "learnmemo.net"
  subject_alternative_names  = ["www.learnmemo.net"]
  validation_method = "DNS"
}

resource "aws_route53_record" "cert_validation" {
  zone_id = data.aws_route53_zone.primary.zone_id
  count = length(aws_acm_certificate.cert.domain_validation_options)
  name = lookup(tolist(aws_acm_certificate.cert.domain_validation_options)[count.index],"resource_record_name")
  type = lookup(tolist(aws_acm_certificate.cert.domain_validation_options)[count.index],"resource_record_type")
  records = [lookup(tolist(aws_acm_certificate.cert.domain_validation_options)[count.index],"resource_record_value")]
  ttl = "300"
}

//resource aws_acm_certificate_validation cert {
//  certificate_arn = aws_acm_certificate.cert.arn
//  validation_record_fqdns = [aws_route53_record.cert_validation[0].fqdn]
//}
