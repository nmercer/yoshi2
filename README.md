# yoshi2
Go GRPC telemetry ingestion service for Raspberry Pi 

## Secrets
```
// TODO: Add these as secrets to each kuber deploy
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
openssl req -new -sha256 -key server.key -out server.csr
openssl x509 -req -sha256 -in server.csr -signkey server.key -out server.crt -days 3650

kubectl create secret generic tls --from-file=server.crt --from-file=server.
key=server.key
```

# minikube
minikube start
minikube stop
minikube ip
minikube service grpc-lb --url

# Docker
// TODO: Reorg these docker image names
docker build -t nmercer88/yoshi2:alpha .
docker push nmercer88/yoshi2:alpha
k rollout restart deployment grpc-server-deployment
