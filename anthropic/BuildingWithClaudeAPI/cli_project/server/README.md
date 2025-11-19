# Notes

## Docker build

To build the Docker image for the server, run the following command from the root directory of the project:

```bash
docker build -t mcp-server --platform linux/amd64 .
```

## Docker run

To run the Docker container for the server, use the following command:

```bash
docker run -p 8000:8000 mcp-server
```
