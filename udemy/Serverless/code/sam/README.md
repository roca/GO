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


npm install -g serverless

sls create -t aws-go-mod -p hello-serverlessA

cd hello-serverless ; sls invoke local -f hello
or
cd hello-serverless ; SLS_DEBUG=* sls invoke local -f hello -d '{\"key\":\"value\"}'

sls deploy --verbose

sls remove
