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
--payload '{"Sessid":"ffa1df92db890ce67c8acc108201e01f2901178e", "Text":"Hello to the lambdas"}' \
outfile

aws lambda invoke \
--function-name Chatter_read \
--payload '{"Sessid":"ffa1df92db890ce67c8acc108201e01f2901178e"}' \
outfile



http://www.desertfoxdev.org