FROM scratch

ENV HOME /var/api
ENV PORT 9090

WORKDIR ${HOME}

# Don't forget to make this binary locally first
# This is a pre-step before build this docker image
# 'CGO_ENABLED=0 GOOS=linux go build -o cloud-native-go'
COPY ./cloud-native-go ${HOME}

EXPOSE ${PORT}

ENTRYPOINT ["./cloud-native-go"]
