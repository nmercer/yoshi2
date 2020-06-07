sh generate_pb.sh

# Create docker image for running migrations
docker build -t nmercer88/yoshi2:migrations -f services/server/migrations/Dockerfile services/server/migrations
docker push nmercer88/yoshi2:migrations

kubectl apply -f k8/

# Create server docker image
docker build -t nmercer88/yoshi2:alpha -f services/server/Dockerfile services/server/
docker push nmercer88/yoshi2:alpha
kubectl rollout restart deployment grpc-server-deployment
