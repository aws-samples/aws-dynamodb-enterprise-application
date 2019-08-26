# Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#  SPDX-License-Identifier: MIT-0

echo "Getting go package dependencies fro serverless function"
echo "******************************"
export GOPATH=$(pwd)
echo "1-Getting AWS SDK"
go get -u github.com/aws/aws-sdk-go
echo "2-Getting Lambda Go"
go get -u github.com/aws/aws-lambda-go/lambda
