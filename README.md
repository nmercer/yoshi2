# yoshi2
Go GRPC telemetry ingestion service and client. This is a project I decided to work on to get more experience with concepts I have been working on professionally.  I wanted to prove to myself that I understood these concepts and could apply them functionally without any support. This is not a real product.

#### GRPC
Data is passed from client <-> server via GRPC. Shared proto files are compiled in the deploy script to both server and client.

#### Docker
Docker images are publicly hosted on docker hub.

#### Kubernetes 
There are a handful of kubernetes features at play in the server. **Minikube** is leveraged here because I had no experience with it and thought it would be interesting to use locally (Most of my experience is with Google Kubernetes Engine). A **LoadBalancer** is used to expose the GRPC server to the world. TLS secrets are mounted via **volumes** and **secrets**. An **initContainers** is used to deploy migrations via a custom [github.com/golang-migrate/migrate](github.com/golang-migrate/migrate) docker image. A **PersistentVolumeClaim** and **PersistentVolume** are used to host the postgres data. The postgres server is running within the cluster for low latency queries.

#### Postgres
Normally to store massive amounts of telemetry data I would not have gone with a relational database. I would have chosen something like Google Cloud Datastore, Google BigQuery or Cassandra. However having worked professionally in NoSQL databases the past few years I wanted to get my hands dirty in SQL again.

#### Golang
Its rad.

# TODO: 
- Document Code
- Unit Testing
- Integration Testing
- CI/CD
- Prometheus
- Go Channels
- Working TLS in Kubernetes
- GRPC Rest Endpoints
- Swagger

# Secrets
```
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
openssl req -new -sha256 -key server.key -out server.csr
openssl x509 -req -sha256 -in server.csr -signkey server.key -out server.crt -days 3650

kubectl create secret generic tls --from-file=server.crt --from-file=server.key=server.key
```

# minikube
minikube start  
minikube stop  
minikube ip  
minikube status  
minikube service grpc-lb --url  

# Deploy
./deploy.sh

# PSQL
brew install golang-migrate  

kubectl exec -it <pod> -- sh  
psql -h localhost -U test --password -p 5432  
create database telemetry;  

# Migrations
// Migration create example  
migrate create -ext sql -dir services/server/migrations/ -seq create_locations_table

# Prometheus
namespace: monitoring
minikube service prometheus-service --url -n monitoring
