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

http://desertfoxdev.org.s3-website-us-east-1.amazonaws.com
http://www.desertfoxdev.org.s3-website-us-east-1.amazonaws.com

arn:aws:s3:::desertfoxdev.org



From with in the chatter directory
    
    Path/to/lambdaup -n new_user

aws lambda invoke \
--function-name Chatter_new_user \
--payload '{"Username":"Matt", "Password":"mattthemaker"}' \
outfile


    Path/to/lambdaup -n login

aws lambda invoke \
--function-name Chatter_login \
--payload '{"Username":"Matt", "Password":"mattthemaker"}' \
outfile

aws lambda update-function-configuration \
--function-name  Chatter_login \
--timeout 4


aws lambda invoke \
--function-name Chatter_say \
--payload '{"Sessid":"2d00632fa3cb8f809909c5bcc5422d4767f46432", "Text":"Hello to the lambdas"}' \
outfile

aws lambda invoke \
--function-name Chatter_read \
--payload '{"Sessid":"2d00632fa3cb8f809909c5bcc5422d4767f46432"}' \
outfile


aws lambda invoke \
--function-name Chatter_envi \
outfile

aws lambda update-function-configuration \
--function-name Chatter_envi \
--environment 'Variables={MyEnv=all}'

aws lambda invoke \
--function-name Logger_logger \
--payload '"Hello people"' \
outfile

http://www.desertfoxdev.org


From udemy:  AWS lambda & serverless course

    https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#input-variable-reference

    http://velocity.apache.org/engine/devel/vtl-reference-guide.html

    http://json-schema.org/learn/examples/address.schema.json

    https://app.swaggerhub.com

    aws lambda add-permission    \
    --function-name "arn:aws:lambda:us-east-1:132172135366:function:eventLogger:test"    \
    --source-arn "arn:aws:execute-api:us-east-1:132172135366:qii3o7ex23/*/GET/"    \
    --principal apigateway.amazonaws.com    --statement-id c9bc8c99-dd58-468b-98ee-32829b4e267c    \
    --action lambda:InvokeFunction

     aws lambda add-permission    \
    --function-name "arn:aws:lambda:us-east-1:132172135366:function:eventLogger:prod"    \
    --source-arn "arn:aws:execute-api:us-east-1:132172135366:qii3o7ex23/*/GET/"    \
    --principal apigateway.amazonaws.com    --statement-id c9bc8c99-dd58-468b-98ee-32829b4e267c    \
    --action lambda:InvokeFunction


    https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-authorization-flow.html