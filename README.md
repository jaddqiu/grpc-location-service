# grpc-location-service

### a simple grpc server  and client for lookup/post(add/update) someone's location

### Usage
1. install the package
```
go get github.com/jaddqiu/grpc-location-service
```

2. start the server

```bash
cd $GOPATH/src/github.com/jaddqiu/grpc-location-service/cmd/locationsvc
go install .
locationsvc
```
3. test the client to see what will happen. os.Args[1] will be the person you want to check, os.Args[2] is the location you want to post for the person.
```bash
cd $GOPATH/src/github.com/jaddqiu/grpc-location-service/cmd/locationcli
go install .
locationcli jadd

locationcli jadd shenzhen
locationcli wy
locationcli wy shenzhen
locationcli wy
```
