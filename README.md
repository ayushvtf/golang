# golang

* Build and run Local

```
go run main.go      
http://localhost:8080
http://localhost:9110
```

* Build and Push Docker image

```
docker build -t golang .
docker tag golang:latest ayushvv/golang-demo:latest
docker push ayushvv/golang-demo:latest      
```


* Deploy via Helm chart 
```
Connect to kubernetes cluster

helm install  helm-golang  helm-golang/ --values  helm-golang/values.yaml

default port: 
8080 for header and 9110 for metric
```

* Service Access

```
header:
kubectl port-forward service/helm-golang-demo 8081:8080 
http://localhost:8081

metric:
kubectl port-forward service/helm-golang-demo 8081:9110 
http://localhost:8081


Total request can be observed from metric 
promhttp_metric_handler_requests_total for status code 200,500 and 503
```

* Custom ports for header and metric servers

```
helm install  helm-golang  helm-golang/ --values  helm-golang/values.yaml --set service.portheader=9091 --set service.portmetric=9092
kubectl port-forward service/helm-golang-demo 8081:9091 
kubectl port-forward service/helm-golang-demo 8081:9092 
```