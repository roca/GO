# [Course Link: Practical intro to Observability](https://www.udemy.com/course/practical-introduction-to-observability)

## SStart Jaeger

[link](https://www.jaegertracing.io/docs/1.54/getting-started/)

```sh
docker run --rm --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.54
```

  visit: [localhost:16686](http://localhost:16686)


## Download promettheus(2.48.0) and Node exporter

- [prometheus-2.48.0.darwin-amd64.tar.gz](https://github.com/prometheus/prometheus/releases/tag/v2.48.0)
- [node_exporter-1.7.0.darwin-amd64.tar.gz]((https://prometheus.io/download/))

## Download and install Grafana

```sh
brew update
brew install grafana

brew services start grafana
```

Go to [Local Browser](localhost:3000) login with admin/admin

## Activating the remote write receiver via '-enable-feature=otlp-writer-receiver' feature flag is deprecated

Use --web.enable-remote-write-receiver instead. 
This feature flag will be ignored in future versions of Prometheus. 

[Code example](https://last9.io/blog/how-to-instrument-golang-app-using-opentelemetry-tutorial-best-practices/)