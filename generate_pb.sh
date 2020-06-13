# TODO: Convert this to a makefile
protoc -I/usr/local/include -I./proto -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:./services/server/telemetry --go_out=plugins=grpc:./services/server/telemetry proto/temperature.proto
protoc -I/usr/local/include -I./proto -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:./services/server/telemetry --go_out=plugins=grpc:./services/server/telemetry proto/location.proto
protoc -I/usr/local/include -I./proto -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis proto/temperature.proto --go_out=plugins=grpc:services/client/telemetry
protoc -I/usr/local/include -I./proto -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis proto/location.proto --go_out=plugins=grpc:services/client/telemetry
