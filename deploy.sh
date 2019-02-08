#!/bin/sh

# Build
GOOS=linux go build ./src/handlers/handler.go
mv ./handler ./src/handlers/

# Create Package
aws cloudformation package \
  --template-file template.yml \
  --output-template-file template-output.yml \
  --s3-bucket handson20180323-sam-packages \
  --profile handson20180323

# Deploy
aws cloudformation deploy \
  --template-file template-output.yml \
  --stack-name appsync-lambda-go \
  --capabilities CAPABILITY_IAM \
  --profile handson20180323
