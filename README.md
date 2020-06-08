# yoshi2
Go GRPC telemetry ingestion service for Raspberry Pi 

## Secrets
```
// TODO: Add these as secrets to each kuber deploy
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
openssl req -new -sha256 -key server.key -out server.csr
openssl x509 -req -sha256 -in server.csr -signkey server.key -out server.crt -days 3650

kubectl create secret generic tls --from-file=server.crt --from-file=server.key=server.key


// TODO: Remove this?
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=grpc.example.com/O=grpc.example.com"
kubectl create secret tls grpc-secret --key tls.key --cert tls.crt
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
// TODO: Description of the way migrations are run

// Migration create example
migrate create -ext sql -dir services/server/migrations/ -seq create_locations_table
