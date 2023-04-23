# lifthus

This is a sample template for lifthus - Below is a brief explanation of what we have generated for you:

```bash
.
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── services                    <-- Microservices
│   ├── lifthus-auth                <-- Auth service
│   └── post                        <-- Post service
└── template.yaml              <-- SAM template
```

## Requirements

- [Docker](https://www.docker.com/community-edition)
- AWS CLI with Administrator permission
- SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
- [Golang](https://golang.org)
- [Node.js](https://nodejs.org/en/)

## Setup process

### Installing dependencies & building the target

1. "swag init" at lifthus-auth
2. "npm install" at lifthus-auth
3. "make build" at lifthus

### Local development

**Invoking function locally through local API Gateway**

```bash
sam local start-api
```

If the previous command ran successfully you should now be able to hit the following local endpoint to invoke your function `http://localhost:3000/`

**Running in native language runtime**

```bash
# lifthus-auth
make go
# post
nest start
```

## Packaging and deployment

Every build and packaging is done by SAM CLI but NestJS must be transpiled to JS before packaging.
Currently this job is done by Makefile. We are migrating to esbuild transpiling which SAM supports natively.

To deploy your application for the first time, run the following in your shell:

```bash
sam deploy --guided
```

The command will package and deploy your application to AWS, with a series of prompts:

- **Stack Name**: The name of the stack to deploy to CloudFormation. This should be unique to your account and region, and a good starting point would be something matching your project name.
- **AWS Region**: The AWS region you want to deploy your app to.
- **Confirm changes before deploy**: If set to yes, any change sets will be shown to you before execution for manual review. If set to no, the AWS SAM CLI will automatically deploy application changes.
- **Allow SAM CLI IAM role creation**: Many AWS SAM templates, including this example, create AWS IAM roles required for the AWS Lambda function(s) included to access AWS services. By default, these are scoped down to minimum required permissions. To deploy an AWS CloudFormation stack which creates or modifies IAM roles, the `CAPABILITY_IAM` value for `capabilities` must be provided. If permission isn't provided through this prompt, to deploy this example you must explicitly pass `--capabilities CAPABILITY_IAM` to the `sam deploy` command.
- **Save arguments to samconfig.toml**: If set to yes, your choices will be saved to a configuration file inside the project, so that in the future you can just re-run `sam deploy` without parameters to deploy changes to your application.

You can find your API Gateway Endpoint URL in the output values displayed after deployment.

### Testing

No tests yet.
