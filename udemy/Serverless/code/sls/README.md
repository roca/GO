npm install -g serverless

sls create -t aws-go-mod -p hello-serverless

cd hello-serverless ; sls invoke local -f hello
or
cd hello-serverless ; SLS_DEBUG=* sls invoke local -f hello -d '{\"key\":\"value\"}'

sls deploy --verbose

sls remove

sls deploy -s prod -f add

npm init

npm install --save-dev serverless-offline
npm install --save-dev serverless-dotenv-plugin

sls offline
sls deploy -s prod

sls logs -f add -s prod --startTime 5m
sls logs -f add -s prod --tail

git push --set-upstream aws-origin dev

npm install --save-dev serverless-domain-manager