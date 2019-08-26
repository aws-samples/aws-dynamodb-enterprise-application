# Building enterprise applications using DynamoDB, Lambda and Go

Demo project for the AWS blog post [Building enterprise applications using DynamoDB, Lambda and Go](https://aws.amazon.com/blogs/database/). It contains the infrastructure script using AWS Cloud Developmeent Script (CDK) along with the code for a lambda function described in the post.

To build the AWS environment, we use **AWS Cloud Development Kit (CDK)**. AWS CDK is an infrastructure as code technology that allows developers to achieve predictable and repeatable deployment by scripting infrastructure definition using **TypeScript, JavaScript, or Python (C#/.Net and Java are also currently in developer preview)**. It is an effective way to deploy an API Gateway infrastructure in that allows the definition of Lambda JSON integration templates directly in the code (as opposed to having escaped JSON code in the middle of a CloudFormation template). To get started with AWS CDK, visit the [Getting Started page](https://docs.aws.amazon.com/cdk/latest/guide/getting_started.html).

The CDK script provided creates a serverless stack with several API endpoints, a Lambda Function, and a DynamoDB table for the purpose of this demo. 

Note the following neat details:

-	We use `envName` variable allowing us to specify the environment name at the moment we call the CDK. This practice allows you to spawn your environment up and down for test purpose without conflict.
```typescript
 const envName = this.node.tryGetContext("envName");
```

-	Giving access to our DynamoDB table to our Lambda Function is as simple as:

```typescript
table.grantReadWriteData(manageCloudrackConfigLambda);
```

-	The code to define our endpoints is lean and clear:

```typescript
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
```

-	CDK using TypeScript allows the definition of complex mapping templates between our API Gateway endpoints and our Lambda function handler. Note that mapping templates are not strictly JSON object but velocity templates. This is why our mapping generator code looks like this

```typescript
  function buildDefaultTpl(functionName: string): string {
    return JSON.stringify({ 
      "userInfo" : buildUserInfoTpl(),
      "request" : "velocity1",
      "subFunction" : functionName,
      "id": "$input.params().path['id']"
    }).replace("\"velocity1\"","$input.json('$$')");
  }
```

We first Stringify a JSON object with placeholders for velocity variables and then replace these variables.

One of the many purposes of the mapping templates described above is to pass a subFunction argument allowing the Lambda function to be aware of which use case to run. Using a single Lambda function as backend for a web administration portal has the following benefits:

1.	It tackles the cold start issue for the Lambda function. Indeed, since the same function is called during the User flow on the front end, the function is always in a “hot” state and provided an improved user experience through more responsiveness.

2.	By checking the code of one single Lambda function in one source control system repository, we avoid having code scattered in hundreds of repositories. We also can reuse shared components easily.

3.	We do not need to manage a large number of Lambda functions for similar or closely related functions. This keeps our infrastructure template lean.

As with most architecture decisions, it also comes with a set of challenges such as a bad code load that could break all use cases at once.  You can mitigate this using Lambda versions in combinations with alias (read here for more info). 

This is a trade-off between operational efficiency and system reliability. Where to set the needle largely depends on your organization’s culture and mechanisms. 

If you choose this approach, you should periodically re-evaluate the split of your function as your model grows to prevent some performance impact and also evaluate the use of Lambda layers.


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

In order to get started you need to have GIT installed and AWS account with CLI access. The install.sh script takes you through the steps to install all relevant componenets needed includind the AWS CDK (you can find the instruction to get this done manually [here](https://docs.aws.amazon.com/cdk/latest/guide/getting_started.html)). This demo runs on MacOSX or Linux due to the Lambda function build script. see the [lambda documentation](https://docs.aws.amazon.com/lambda/latest/dg/lambda-go-how-to-create-deployment-package.html) for the windows build guide.

You also need to have installed the AWS GO SDK and be able to compile GO code locally

### Installing

**If you are not interested in using the CDK, use the pregenerated [cloudFormation template](./cloud-infra-demo/crud-demo-stack.json). You can find more information on the [AWS Documentation](https://aws.amazon.com/cloudformation/getting-started/)**

#### Installing

If you start from a clean EC2 instance with Amazon Linux (or any linux distribution using the yum package manager), follow the steps below to run the demo.

```shell
git clone xxxx
cd aws-dynamodb-crud-demo
sh install.sh
sh deploy.sh <environementName>
```

Else, you can skip the `install.sh` process and install manually the following componenst on your platform:
1. Make sure your AWS CLI is installed and configured correctly. (for more information, go to the [AWS documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html))
```shell
aws configure
```
2. Download and Install [NodeJS] (https://nodejs.org/en/download/)
3. Install [Golang](https://golang.org/doc/install)
4. Install the AWS CDK globally. (check [here](https://docs.aws.amazon.com/cdk/latest/guide/getting_started.html) for more details
```shell
npm i -g aws-cdk
```
4. Setup the infrastructure script
```shell
cd cloud-infra-demo
npm install
```
5. Get the AWS go SDK packages
```shell
cd ../lambda
sh env.sh
```
6. Once you environement is setup

this will deploy a full AWS stack including:
* A set of API enpoints on AWS API Gateway
* A lambda function interaction with DynamoDB using the AWS GO SDK
* A DynamoDB table

Each resource is unique to for the provided environment name so that you  can deploy multiple stacks in parralel.

## Running the tests

The command above will output the an API gateway enpoint url that can be used for testing the rest API. Something like:
```
Outputs:
CloudInfraDemoStack.awsCrudDemoEndpointE705E337 = https://<random_sequence>.execute-api.<region>.amazonaws.com/prod/
```

You can then test the application uing the following script

```shell
curl -d '{"name" : "My First hotel","description" : "This is a great property"}' -H "Content-Type: application/json" -X POST https://<random_sequence>.execute-api.<region>.amazonaws.com/prod/config
```

Then use the ID provided in the response for the subsequent calls

```shell
ID=<id_fom-create_response>
curl -d '{"code" : "'$ID'","roomTypes": [{"code" : "DBL","name" : "Room with Double Beds","description" : "Stadard Room Type with double bed"}]}' -H "Content-Type: application/json" -X POST https://<random_sequence>.execute-api.<region>.amazonaws.com/prod/config/$ID/rooms/type
curl https://<random_sequence>.execute-api.<region>.amazonaws.com/prod/config/$ID
```

## Built With

* [AWS](https://aws.amazon.com) - Build on AWS
* [GO](https://golang.org/) - Build using teh Goland AWS SDK

## Contributing

Please read [CONTRIBUTING.md](to be added) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Geoffroy Rollat** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This sample code is made available under the MIT-0 license. See the LICENSE file.