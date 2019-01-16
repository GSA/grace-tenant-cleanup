output "aws_secret_access_key" {
  value = "${aws_iam_access_key.test_deployer.secret}"
}

output "aws_access_key_id" {
  value = "${aws_iam_access_key.test_deployer.id}"
}
