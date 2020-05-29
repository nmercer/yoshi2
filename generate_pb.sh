# TODO: 
# Possibly generate all protos with one command?
# Possibly have that shared?
protoc -I proto/ proto/temperature.proto --go_out=plugins=grpc:services/server/telemetry
protoc -I proto/ proto/location.proto --go_out=plugins=grpc:services/server/telemetry
protoc -I proto/ proto/temperature.proto --go_out=plugins=grpc:services/client/telemetry
protoc -I proto/ proto/location.proto --go_out=plugins=grpc:services/client/telemetry
