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

### Cobra

- Cobra is a library for creating powerful modern CLI applications.

[link](https://github.com/spf13/cobra)