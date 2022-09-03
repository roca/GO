# Commands

- minikube start --driver=hyperkit -p singlenode-demo
- kubectl port-forward svc/github-deploy EXTERNAL-PORT:8080
- ngrok http EXTERNAL-PORT

Now configure your webhook using the out put from ngrok
