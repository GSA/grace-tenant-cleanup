# Archived #

We're looking at using [aws-nuke](https://github.com/rebuy-de/aws-nuke) instead.

# grace-tenant-cleanup [![CircleCI](https://circleci.com/gh/GSA/grace-tenant-cleanup.svg?style=svg)](https://circleci.com/gh/GSA/grace-tenant-cleanup)

Go program to delete services from AWS accounts prior to decommissioning

## Supported Services ##

- EC2 Instances

## Usage Instructions ##

**WARNING There are no guardrails.  This program will delete all supported services
with no undo or option to quit**

### Prerequisites ###

- An AWS account in which you have access to Get, List, Describe, Terminate,
Deregister and Delete the supported services.

- Properly configured credentials for access to the AWS account.  Recommend
using [Configuration and Credential Files](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)

- If you are using configuration and credential files, `export AWS_SDK_LOAD_CONFIG=1`

- If you want to restrict the program to specific regions, provide a comma separated
list in the `regions` environment variable.  For example: `export regions="us-east-1,us-west-1,us-west-2"`

- The [Go language compiler](https://golang.org/doc/install)

- Steps to build:

    1. `go get github.com/GSA/grace-tenant-cleanup`

- Steps to clean up the account:

    1. `grace-tenant-cleanup`

## Public domain ##

This project is in the worldwide [public domain](LICENSE.md). As stated in [CONTRIBUTING](CONTRIBUTING.md):

> This project is in the public domain within the United States, and copyright and related rights in the work worldwide are waived through the [CC0 1.0 Universal public domain dedication](https://creativecommons.org/publicdomain/zero/1.0/).
>
> All contributions to this project will be released under the CC0 dedication. By submitting a pull request, you are agreeing to comply with this waiver of copyright interest.
