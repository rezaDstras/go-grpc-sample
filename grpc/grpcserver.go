package user_grpc

import (
	"context"
	"github.com/rezaDastrs/protocolBuffer/datalayer"
)

type GrpcServer struct {
	dbHandler                      *datalayer.SqlHandler
	UnimplementedUserServiceServer *unimplementedUserServiceServer
}

type unimplementedUserServiceServer struct {

}

func (server *GrpcServer) mustEmbedUnimplementedUserServiceServer() {
	server.UnimplementedUserServiceServer = &unimplementedUserServiceServer{}
}

func NewGrpcServer(connString string) (*GrpcServer, error) {
	db, err := datalayer.CreateConnection(connString)
	if err != nil {
		return nil, err
	}
	return &GrpcServer{
		dbHandler: db,
	}, err
}

//implement the grpc user server methods in user_grpc.pb.go


func (server *GrpcServer) GetUser(ctx context.Context, r *Request) (*User, error) {
	user, err := server.dbHandler.GetUserByName(r.GetName())
	if err != nil {
		return nil, err
	}
	return convertToGrpcUser(user), nil
}
func (server *GrpcServer) GetAllUsers(r *Request, stream UserService_GetAllUsersServer) error  {
	users, err := server.dbHandler.GetAllUsers()
	if err != nil {
		return err
	}
	for _, user := range users {
		err := stream.Send(convertToGrpcUser(user))
		if err != nil {
			return err
		}
	}
	return nil
}

func convertToGrpcUser(user datalayer.User) *User {
	return &User{
		Id : int32(user.Id),
		Name : user.Name,
		Family : user.Family,
	}
}