env=$1
# Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#  SPDX-License-Identifier: MIT-0x

echo "**********************************************"
echo "* Deploying DynamoDB CRUD demo '$env' "
echo "***********************************************"
if [ -z "$env" ]
then
    echo "Environment Must not be Empty"
    echo "Usage:"
    echo "sh deploy.sh <env>"
else
    
    echo "0-Making npm utility available (hacky: fix this)"
    export NVM_DIR="$HOME/.nvm"
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
    [ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion 
    echo "1-Building Lambda function for environment '$env' using AWS-GO-SDK"    
    cd lambda
    sh lbuild.sh $env
    echo "2-Deploying Stack for environement '$env' using AWS CDK"
    cd ../cloud-infra-demo
    sh deploy.sh $env
fi