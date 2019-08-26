
# Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#  SPDX-License-Identifier: MIT-0

env=$1
echo "**********************************************"
echo "* Deploying DynamoDB CRUD demo stack environment '$env' "
echo "***********************************************"
if [ -z "$env" ]
then
    echo "Environment Must not be Empty"
    echo "Usage:"
    echo "sh deploy.sh <env>"
else
    echo "0-Building stack for environement '$env' "
    npm run build
    echo "1-Synthesizing CloudFormation template for environement '$env' "
    cdk synth -c envName=$env > crud-demo-stack.json
    echo "2-Analayzing changes for environement '$env' "
    cdk diff -c envName=$env
    echo "2-Deploying infrastructure for environement '$env' "
    cdk deploy -c envName=$env
fi