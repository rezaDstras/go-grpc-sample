package main

import (
	"context"
	"flag"
	"fmt"
	usergrpc "github.com/rezaDastrs/protocolBuffer/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	op := flag.String("op", "s", "s for Server and c for Client.")
	flag.Parse()

	switch strings.ToLower(*op) {
	case "s":
		runGrpcServer()
	case "c":
		runGrpcClient()
	}
}

func runGrpcServer() {
	grpclog.Println("starting server ...")
	listen, err := net.Listen("tcp", ":8282")
	if err != nil {
		log.Fatalln("failed to listen")
	}
	grpclog.Println("listening on localhost:8282")

	var option []grpc.ServerOption
	server := grpc.NewServer(option...)
	userServer, err := usergrpc.NewGrpcServer("root:@tcp(127.0.0.1:3306)/bookstore")
	if err != nil {
		log.Fatalln(err)
	}

	usergrpc.RegisterUserServiceServer(server, userServer)

	err = server.Serve(listen)
	if err != nil {
		log.Fatalln(err)
	}
}

func runGrpcClient(){
	//connect to grpc server
	conn, err := grpc.Dial("127.0.0.1:8282",grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	//up server client
	client := usergrpc.NewUserServiceClient(conn)

	input:=""
	fmt.Println("all users ? (y/n)")
	fmt.Scanln(&input)

	//check situation
	if strings.EqualFold(input,"y"){
		//get all users
		users, err := client.GetAllUsers(context.Background(), &usergrpc.Request{})
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(users)
		//unlimited for => for stream
		for  {
			user, err := users.Recv()
			//if data is empty or done (empty or full)
			if err == io.EOF {
				break
			}
			if err !=nil {
				log.Fatalln(err)
			}
			fmt.Println(user)
		}
		return
	}

	fmt.Println("name?")
	fmt.Scanln(&input)

	user, err := client.GetUser(context.Background(),&usergrpc.Request{Name: input})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user)

}



