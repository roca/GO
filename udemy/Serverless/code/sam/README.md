aws cloudformation package \
--template-file template.yaml \
--output-template-file sam-template.yaml \
--s3-bucket romel-sam

aws s3 mb s3:://romel-sam

aws cloudformation deploy \
--template-file sam-template.yaml \
--stack-name hello-sam-stack \
--capabilities CAPABILITY_IAM


sam package \
--template-file template.yaml \
--output-template-file sam-template.yaml \
--s3-bucket romel-sam

sam deploy \
--template-file sam-template.yaml \
--stack-name sam-rest-api-stack \
--capabilities CAPABILITY_IAM


sam logs -n GetUser --stack-name sam-rest-api-stack --tail
