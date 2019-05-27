Save time and resources by building applications with Golang and deploying them to AWS Lambda

API; endpoint, Invoke URL: https://kl4mnqn3bc.execute-api.us-east-1.amazonaws.com/Prod



aws lambda create-function \
--function-name questionresponse \
--runtime go1.x \
--role arn:aws:iam::132172135366:role/service-role/donothing \
--handler main \
--zip-file fileb://main.zip

aws lambda invoke \
--function-name questionresponse \
--payload '{"Question": "How tall are you?"}'