# Source files and useful information

## All the source files can be found at

- https://github.com/wardviaene/golang-for-devops-course

## Visual Studio Code can be downloaded from

- https://code.visualstudio.com/download

## If necessary Golang can be downloaded separately from

- https://go.dev/dl/

## Golang official website

- https://go.dev/

### Run Jenkins in Docker

``` sh
docker run -v jenkins:/var/jenkins_home -p 8090:8080 -it --rm --name jenkins jenkins/jenkins:lts
```

### OpenID LDAP Wiki

[Openid-configuration](https://ldapwiki.com/wiki/Openid-configuration)


aws sts assume-role-with-web-identity \
 --role-arn $ROLE_ARN \
--web-identity-token $TOKEN \
--role-session-name edward

AWS_ACCESS_KEY_ID=ASIAR5RQSXPDGVRF7I6N \
AWS_SECRET_ACCESS_KEY=1QtVFpzOlLnaMSe57rBb0P5MViWrpZkqK1AyIrqh \
AWS_SESSION_TOKEN=FwoGZXIvYXdzEN7//////////wEaDP2YBVfzuLuLqophfiKnAuW/i/oUnD8f4BK507R/7CdSIb0r4ExEorpAMd6hm8NCaze3/FTPm+yxShi4iMKXDEVUwf4zsTX4V+iuSYwRr5GVUEY0VdgRq00A85UKwZU0icD9Rewa904hAT8tm6X2XjEhGl/nh3r7oydmrgL9iX9cfKV+U0R98VE/ZL7njroSSwu5Q4Mw9sTjm26n0T9DnEhxQ1BeadHYTDDP6k3NZWsbnjgE1WvHcWCBqIOu7+Mn2LuqHdEmuSN9UjT/xO/F4vHaO4XlpfKPSjKzFT/LVk+8i3qzqeVjTODokgMHon8A94DII4y6uCZMxsooLjT3AuNLiIiZ6fm6anUoA/QBeBA4SOlbboBX4R2xRIwteXf2YtTUeQu8+2K+Ufo68mW5p2LA4Oyj5/co9sucmgYyJGzb5J9H/06fsNPpcBL2qaycY6T0W6zTslNG2O8uXdgSkDgSpA== \
aws s3 ls