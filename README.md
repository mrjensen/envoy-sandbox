#Envoy demo/sandbox

##### Requirements
- Docker Engine



```bash
# Bring up cluster with 3 replicas of service1
docker-compose up --build -d
docker-compose scale service1=3

# Open Grafana dashboard at http://localhost:3000 - login: admin/foobar
curl http://localhost:8000/service/1
curl http://localhost:8000/service/2

# Demo 1: Make one of the service1 replica return 503 temporarily
curl http://localhost:8000/service/1/sleep
# See passive health check in "Cluster membership" 


# Demo 2: Inject status code 400 errors in service1
# Change "static_resources.listeners[0].filter_chains[0].filters[0].config.http_filters[0].config.abort.percentage.numerator" in service1/service-envoy.yaml to the percentage of errors you would like.
docker-compose up -d --force-recreate --build service1

# FYI: --force-recreate implicitly also scales to 1 replica of service1 
```


##### Containers exposed on localhost ports 
- [envoy admin panel :8001](http://localhost:8001/ )
- [front-envoy :8000](http://localhost:8000)
- [grafana admin/foobar :3000](http://localhost:3000) 
- [cAdvisor :8080](http://localhost:8080)
- [prometheus :9090](http://localhost:9090)
- [node exporter :9100](http://localhost:9100)

##### Containers not exposed on localhost ports
- service1 (:8080)
- service2 (:8080)

#Credit
This is heavily inspired by envoy proxy sandbox from the [envoy docs](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/front_proxy.html)


[Envoy dashboard](https://grafana.com/dashboards/7250)


[Prometheus demo project](https://github.com/vegasbrianc/prometheus)