# Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#  SPDX-License-Identifier: MIT-0

echo "0-Setup or validate your AWS credentials and region"
aws configure
echo "1-Installing GO"
sudo yum install go
echo "2-Installing nodejs"
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.34.0/install.sh | bash
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion    
nvm install 10.16.2
echo "3-installing cdk globally"
npm i -g aws-cdk
echo "4-Setup infrastructure script"
cd cloud-infra-demo
npm install
echo "5-Setup lamda function"
cd ../lambda
sh env.sh