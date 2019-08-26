#!/bin/sh
# Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: MIT-0
env=$1
echo "**********************************************"
echo "* Building Lambda for DynamoDB Demo '$env' "
echo "***********************************************"
if [ -z "$env" ]
then
    echo "Environment Must not be Empty"
    echo "Usage:"
    echo "sh lbuild.sh <env>"
else
echo "1-Setting up environement"
export GOPATH=$(pwd)
echo "GOPATH set to $GOPATH"
echo "2-Cleaning old builds"
mkdir src/main
export GOOS=linux
rm main main.zip
echo "3-Building application"
go build -o main src/main/main.go
zip main.zip main
echo "4-Deploying to Lambda"
sh push.sh $env
fi

