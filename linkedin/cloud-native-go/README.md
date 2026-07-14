cloud-native-go course

## Docker
### Remove dangling images

    - docker rmi -f $(docker images -f "dangling=true" -q)

## Kubernetes

 - kubectl port-forward cloud-native-go 9090:9090
 
