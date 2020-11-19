# grpc
Sample to use grpc

# Learning Note

## Prepare

### Install gRPC
```
go get -u google.golang.org/grpc
```

### Install protoc
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Development

### Create a .proto file
**pb.proto**
```
syntax = "proto3";

//option go_package = "pb";

service LocalService {
    rpc BrowseForFolder(BrowseInfo) returns (Folder) {}
    rpc Alive(Empty) returns (AliveResponse) {}
}

//Empty -
message Empty {
}

//BrowseInfo -
message BrowseInfo {
    string title = 1;
    string initialdir = 2;
}

//Folder - BrowseForFolder response
message Folder {
    string path = 1;
}

//AliveResponse - alive response
message AliveResponse {
    int32 processid = 1;
    string homedir = 2;
    string accesskey = 3;
    string secretkey = 4;
    string region = 5;
}
```
#### API
1. BrowseForFolder(BrowseInfo)
1. Alive(Empty)

### Generate go file
```
protoc --go_out=plugins=grpc:pb pb.proto
```

## Server
**main.go**
```
	//new server
	server := grpc.NewServer()
	svc := &service.LocalService{}
	go func() {
		//resigter
		pb.RegisterLocalServiceServer(server, svc)
		fmt.Printf("listening on %v\n", listenon)
		server.Serve(port)
	}()
```

## Client
**pb/client.go**
```
//GetLocalService -
func GetLocalService(port int) (*LocalService, error)
```

### caller
Ext. **Alive**
```
    ls, err := pb.GetLocalService(cfg.Port)
    if err != nil {
      logger.Tracef("GetLocalService %v", err)
      srv._uikit = false
      srv._uikitpid = -1
      continue
    }
    //call alive
    res, err := ls.RPC().Alive(context.Background(), &pb.Empty{})
    if err != nil {
    }
```

# so simple. right? :)
