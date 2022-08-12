# go-grpc-sample



follow up : 

#Portocol buffer

* Download protocol buffer compiler => with home-brew in  google.github.io/proto-lens/installing-protoc.html
* Install plugin in project to translate protocol buffer to language => go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
* Create file.proto file 
* Compile that with command => protoc --go_out=. file.proto




#Grpc

* Install grip compiler => https://grpc.io/docs/languages/go/quickstart/
* Run command protoc --go_out=. --go_opt=paths=source_relative \
*     --go-grpc_out=. --go-grpc_opt=paths=source_relative \
*     file.proto 
 
