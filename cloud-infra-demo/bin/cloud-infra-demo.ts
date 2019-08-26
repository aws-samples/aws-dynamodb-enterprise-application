#!/usr/bin/env node
import 'source-map-support/register';
import cdk = require('@aws-cdk/core');
import { CloudInfraDemoStack } from '../lib/cloud-infra-demo-stack';

const app = new cdk.App();
const envName = app.node.tryGetContext("envName");
new CloudInfraDemoStack(app, 'CloudInfraDemoStack' + envName);
