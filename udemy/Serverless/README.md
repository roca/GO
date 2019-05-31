Save time and resources by building applications with Golang and deploying them to AWS Lambda

API; endpoint, Invoke URL: https://kl4mnqn3bc.execute-api.us-east-1.amazonaws.com/Prod

Examples create/invoke:

aws lambda create-function \
--function-name questionresponse \
--runtime go1.x \
--role arn:aws:iam::132172135366:role/service-role/donothing \
--handler main \
--zip-file fileb://main.zip

aws lambda invoke \
--function-name questionresponse \
--payload '{"Question": "How tall are you?"}' \
outfile


aws lambda invoke \
--function-name Lambda_UP_1_firstgo \
--payload '{"Username":"Romel"}' \
outfile

aws lambda invoke \
--function-name Chatter_servicewriter \
--payload '{"Txt":"Pete"}' \
outfile

http://minisite.storyfeet.com.s3-website-us-east-1.amazonaws.com