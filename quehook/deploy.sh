#!/bin/bash

GOARCH=amd64 GOOS=linux go build -ldflags "-X main.HANDLER=CREATE" -o create
zip quehook-create.zip create
aws lambda update-function-code --function-name quehook-query-create --zip-file fileb://quehook-create.zip --region us-east-1

GOARCH=amd64 GOOS=linux go build -ldflags "-X main.HANDLER=RUN" -o run
zip quehook-run.zip run
aws lambda update-function-code --function-name quehook-query-run --zip-file fileb://quehook-run.zip --region us-east-1

GOARCH=amd64 GOOS=linux go build -ldflags "-X main.HANDLER=DELETE" -o delete
zip quehook-delete.zip delete
aws lambda update-function-code --function-name quehook-query-delete --zip-file fileb://quehook-delete.zip --region us-east-1

GOARCH=amd64 GOOS=linux go build -ldflags "-X main.HANDLER=SUBSCRIBE" -o subscribe
zip quehook-subscribe.zip subscribe
aws lambda update-function-code --function-name quehook-subscription-subscribe --zip-file fileb://quehook-subscribe.zip --region us-east-1

GOARCH=amd64 GOOS=linux go build -ldflags "-X main.HANDLER=UNSUBSCRIBE" -o unsubscribe
zip quehook-unsubscribe.zip unsubscribe
aws lambda update-function-code --function-name quehook-subscription-unsubscribe --zip-file fileb://quehook-unsubscribe.zip --region us-east-1
