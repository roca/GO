cloud-native-go course

## Docker
### Remove dangling images

    - docker rmi -f $(docker images -f "dangling=true" -q)
