

kubectl apply -f k8/

docker build -t nmercer88/yoshi2:alpha -f services/server/Dockerfile services/server/
docker push nmercer88/yoshi2:alpha
kubectl rollout restart deployment grpc-server-deployment