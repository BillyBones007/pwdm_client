package grpcclient

import (
	"context"

	"github.com/BillyBones007/pwdm_client/internal/datatypes"
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
	"github.com/BillyBones007/pwdm_client/internal/tools/encrypttools"
	pb "github.com/BillyBones007/pwdm_service_api/api"
	"google.golang.org/grpc"
)

// Registration - registers a new user.
func (c *ClientGRPC) Registration(user models.UserModel) (authUser string, err error) {
	conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds))
	if err != nil {
		return authUser, err
	}
	defer conn.Close()
	client := pb.NewAuthServiceClient(conn)
	req := &pb.AuthReq{Login: user.Login, Password: user.Password}
	resp, err := client.Create(context.Background(), req)
	if err != nil {
		return authUser, err
	}
	c.AuthFlag = true
	c.Token = resp.Token
	c.Encrypter = encrypttools.NewEncrypter(user.Password)
	authUser = user.Login

	return authUser, nil
}

// Authentication - user authentication.
func (c *ClientGRPC) LogIn(user models.UserModel) (authUser string, err error) {
	conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds))
	if err != nil {
		return authUser, err
	}
	defer conn.Close()
	client := pb.NewAuthServiceClient(conn)
	req := &pb.AuthReq{Login: user.Login, Password: user.Password}
	resp, err := client.Enter(context.Background(), req)
	if err != nil {
		return authUser, err
	}
	c.AuthFlag = true
	c.Token = resp.Token
	c.Encrypter = encrypttools.NewEncrypter(user.Password)
	authUser = user.Login

	return authUser, nil
}

// SignOut - user logout.
func (c *ClientGRPC) SignOut() {
	c.AuthFlag = false
	c.Token = ""
	c.Encrypter = encrypttools.NewEncrypter("")
	c.Config.Storage.Clear(datatypes.LoginPasswordDataType, datatypes.CardDataType, datatypes.TextDataType, datatypes.BinaryDataType)
}
