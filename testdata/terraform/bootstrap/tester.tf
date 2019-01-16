provider "aws" {
  alias = "test"

  assume_role {
    role_arn = "arn:aws:iam::${var.test_account_id}:role/${var.iam_role_name}"
  }
}

resource "aws_iam_access_key" "test_deployer" {
  provider = "aws.test"
  user     = "${aws_iam_user.test_deployer.name}"
}

resource "aws_iam_user" "test_deployer" {
  provider = "aws.test"
  name     = "circle-test-deployer"
}

resource "aws_iam_user_policy_attachment" "test_deployer_attach" {
  provider = "aws.test"
  user     = "${aws_iam_user.test_deployer.name}"

  # AWS-managed policy
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
}
