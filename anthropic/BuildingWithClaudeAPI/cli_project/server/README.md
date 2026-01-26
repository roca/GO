# Notes

## Docker build

To build the Docker image for the server, run the following command from the root directory of the project:

```bash
docker build -t mcp-server --platform linux/amd64 .
```

## Docker run

To run the Docker container for the server, use the following command:

```bash
docker run -p 8080:8080 mcp-server
```

## Docker tag for AWS ECR

To tag the Docker image for AWS ECR, use the following command:

```bash
docker tag mcp-server:latest 154919775133.dkr.ecr.us-east-1.amazonaws.com/app_dev/general:mcp-server
```

## Docker push to AWS ECR

To push the Docker image to AWS ECR, use the following command:

```bash
docker push 154919775133.dkr.ecr.us-east-1.amazonaws.com/app_dev/general:mcp-server
```
