# grpc-location-service

### a simple grpc server  and client for lookup/post(add/update) someone's location

### Usage
1. install the package
```
go get github.com/jaddqiu/grpc-location/service
```

2. start the server

```bash
cd $GOPATH/src/github.com/jaddqiu/grpc-location-service/cmd/locationsvc
go install .
locationsvc
```
3. test the client to see what will happen
```bash
cd $GOPATH/src/github.com/jaddqiu/grpc-location-service/cmd/locationsvc
go install .
locationsvc jadd

locationsvc jadd shenzhen
locationsvc wy
locationsvc wy shenzhen
locationsvc wy
```
