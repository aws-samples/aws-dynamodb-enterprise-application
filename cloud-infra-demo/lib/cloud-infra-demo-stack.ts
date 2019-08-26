// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

import cdk = require('@aws-cdk/core');
import lambda = require('@aws-cdk/aws-lambda');
import apigateway = require('@aws-cdk/aws-apigateway');
import dynamodb = require('@aws-cdk/aws-dynamodb');

/////////////////////////////////////////////////////////////////////
//Infrastructure Script for the AWS DynamoDB Blog Demo
//This script takes the name of the Environement to spawn and create 
// A lambda function, A DynamoDB table and a set of Api Gateway enpoints.
/////////////////////////////////////////////////////////////////////
export class CloudInfraDemoStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    
    super(scope, id, props);
    
    const envName = this.node.tryGetContext("envName");
    if(!envName) {
      throw new Error('No environemnt name provided for stack');
    }
    /*********************************
     * LAMBDA FUNCTIONS DECLARATIONS
     * **********************************/
    
    const manageCloudrackConfigLambda = new lambda.Function(this, 'awsDynamoCrudDemo'+envName, {
      code: new lambda.AssetCode('../lambda/main.zip'),
      functionName: "manageCloudrackConfig"+envName,
      handler: 'main',
      runtime: lambda.Runtime.GO_1_X,
      environment: {
        LAMBDA_ENV: envName
      }
    });

    /**********************
    * DYNAMO DB TABLE
    *************************/
   const table = new dynamodb.Table(this, 'Table', {
    tableName: "aws-crud-demo-config-" + envName,
    partitionKey: { name: 'code', type: dynamodb.AttributeType.STRING },
    sortKey: {name: 'itemType', type: dynamodb.AttributeType.STRING }
    });

    //grant lamda fnuction access to our dynamo table
   table.grantReadWriteData(manageCloudrackConfigLambda);

  /**********************
   * API GATEWAY
   *************************/
  const api = new apigateway.RestApi(this, 'awsCrudDemo'+envName, {
    restApiName: 'awsCrudDemo'+envName,
    deployOptions: {
      loggingLevel: apigateway.MethodLoggingLevel.INFO,
      dataTraceEnabled: true
    }
  });

  const config = api.root.addResource('config',{
    defaultMethodOptions: {
      methodResponses: [
        {statusCode: "200"}
      ]
    }
  });

  addCorsOptions(config);
  config.addMethod('GET', buildConfigLambdaIntegration("listHotels"),methodOptions());
  config.addMethod('POST', buildConfigLambdaIntegration("saveHotel"),methodOptions());
  //config/{id}
  const configId = config.addResource('{id}');
  configId.addMethod('GET', buildConfigLambdaIntegration("getHotel"),methodOptions());
  addCorsOptions(configId);
  //config/{id}/rooms
  const configIdRooms= configId.addResource('rooms');
  addCorsOptions(configIdRooms);
  //config/{id}/rooms/type
  const configIdRoomsType= configIdRooms.addResource('type');
  addCorsOptions(configIdRoomsType);
  configIdRoomsType.addMethod('POST', buildConfigLambdaIntegration("addRoomType"),methodOptions());
  //config/{id}/rooms/type/{typecode}
  const configIdRoomsTypeTypeCode= configIdRoomsType.addResource('{typecode}');
  addCorsOptions(configIdRoomsTypeTypeCode);
  configIdRoomsTypeTypeCode.addMethod('DELETE', buildConfigLambdaIntegration("deleteRoomType",buildDeleteRoomTypeTpl("deleteRoomType")),methodOptions());

  /********************
   * Utilility methods
   * This is a set of reuable methods to create some standard configuration objects
   * to be used for the definition of API gateway enpoints (this includes CORS configuration,
   * Lambda integartion mapping template...)
   */

  function methodOptions(): apigateway.MethodOptions {
    return {
      methodResponses: [
        {statusCode: "200"}
      ]
    } 
  }

  function buildConfigLambdaIntegration(fnName : string, customBody?: string): any{
    return new apigateway.LambdaIntegration(manageCloudrackConfigLambda,{
      proxy: false,
      requestTemplates: {
        'application/json': customBody || buildDefaultTpl(fnName)
      },
      integrationResponses : [{
        statusCode: "200"
      }]
    });
  }

  function buildDefaultTpl(functionName: string): string {
    return JSON.stringify({ 
      "userInfo" : buildUserInfoTpl(),
      "request" : "velocity1",
      "subFunction" : functionName,
      "id": "$input.params().path['id']"
    }).replace("\"velocity1\"","$input.json('$$')");
  }

  function buildUserInfoTpl(): any {
    return {
      "sub" : "$context.authorizer.claims.sub",
      "email" : "$context.authorizer.claims.email",
      "username" : "$context.authorizer.claims['cognito:username']"
    }
  }

  function buildDeleteRoomTypeTpl(functionName: string): string {
    return JSON.stringify({ 
      "userInfo" : buildUserInfoTpl(),
      "request" : {
        "code" : "$input.params().path['id']",
        "tags" : [
            {
                "code" : "$input.params().path['tagId']"
            }
        ]
    },
      "subFunction" : functionName
    });
  }

  function addCorsOptions(apiResource: apigateway.IResource) {
    apiResource.addMethod('OPTIONS', new apigateway.MockIntegration({
      integrationResponses: [{
        statusCode: '200',
        responseParameters: {
          'method.response.header.Access-Control-Allow-Headers': "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent'",
          'method.response.header.Access-Control-Allow-Origin': "'*'",
          'method.response.header.Access-Control-Allow-Credentials': "'false'",
          'method.response.header.Access-Control-Allow-Methods': "'OPTIONS,GET,PUT,POST,DELETE'",
        },
      }],
      passthroughBehavior: apigateway.PassthroughBehavior.NEVER,
      requestTemplates: {
        "application/json": "{\"statusCode\": 200}"
      },
    }), {
      methodResponses: [{
        statusCode: '200',
        responseParameters: {
          'method.response.header.Access-Control-Allow-Headers': true,
          'method.response.header.Access-Control-Allow-Methods': true,
          'method.response.header.Access-Control-Allow-Credentials': true,
          'method.response.header.Access-Control-Allow-Origin': true,
        },  
      }]
    })
  }
    
  }
}
