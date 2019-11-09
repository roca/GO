npm install -g serverless

sls create -t aws-go-mod -p hello-serverless

cd hello-serverless ; sls invoke local -f hello
or
cd hello-serverless ; SLS_DEBUG=* sls invoke local -f hello -d '{\"key\":\"value\"}'

sls deploy --verbose

sls remove
