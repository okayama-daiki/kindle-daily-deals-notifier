#!/bin/bash

function_name="notifyKindleDailyDeals"

GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go
zip myFunction.zip bootstrap

if aws lambda list-functions | grep -q $function_name; then
  aws lambda update-function-code \
    --function-name $function_name \
    --zip-file fileb://myFunction.zip
else
  aws lambda create-function \
    --function-name $function_name \
    --runtime provided.al2023 \
    --architectures arm64 \
    --role arn:aws:iam::828951707561:role/awslambdaBasicExecutionRole \
    --environment Variables="{LINE_CHANNEL_ACCESS_TOKEN=YOUR_CHANNEL_ACCESS_TOKEN,LINE_TARGET_ID=YOUR_TARGET_ID}" \
    --handler bootstrap \
    --zip-file fileb://myFunction.zip
fi
