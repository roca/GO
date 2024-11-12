# [Ultimate Go Software Design With Kubernetes 2](https://courses.ardanlabs.com/courses/take/ultimate-go-software-design-with-kubernetes-2/lessons/57299253-18-1-check-for-dependency-upgrades-rebuild-run)

@00:00 Module 18.1

- [code](https://github.com/ardanlabs/service.git)


## READ

- [Kubernetes Memory Limits and Go](https://www.ardanlabs.com/blog/2024/02/kubernetes-memory-limits-go.html)[:->](https://www.ardanlabs.com/blog/2024/02/kubernetes-memory-limits-go.html)

- [Kubernetes CPU Limits and Go](https://www.ardanlabs.com/blog/2024/02/kubernetes-cpu-limits-go.html)[:->](https://www.ardanlabs.com/blog/2024/02/kubernetes-cpu-limits-go.html)


## Watch

- [GopherCon SG (2023) - K8s CPU Limits and Go](https://www.youtube.com/watch?v=Dm7yuoYTx54&list=PLq2Nv-Sh8Eba2gEaId35K2aAUFdpbKx9D&index=6&themeRefresh=1)[:->](https://www.youtube.com/watch?v=Dm7yuoYTx54&list=PLq2Nv-Sh8Eba2gEaId35K2aAUFdpbKx9D&index=6&themeRefresh=1)


## Layers

api > app > business > storage

## Push to AWS ECR

```bash
docker tag localhost/ardanlabs/auth:0.0.1 132172135366.dkr.ecr.us-east-1.amazonaws.com/ardanlabs/auth:latest 
docker push 132172135366.dkr.ecr.us-east-1.amazonaws.com/ardanlabs/auth:latest

docker tag localhost/ardanlabs/sales:0.0.1 132172135366.dkr.ecr.us-east-1.amazonaws.com/ardanlabs/sales:latest
docker push 132172135366.dkr.ecr.us-east-1.amazonaws.com/ardanlabs/sales:latest
```