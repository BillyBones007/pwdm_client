package grpcclient

import (
	"github.com/BillyBones007/pwdm_client/internal/datatypes"
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
	pb "github.com/BillyBones007/pwdm_service_api/api"
	"google.golang.org/grpc"
)

// SendLogPwd - send login and password to the server.
func (c *ClientGRPC) SendLogPwd(data models.LogPwdModel) (string, error) {
	ctx := c.getContext()
	conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds), grpc.WithUnaryInterceptor(c.SendLogPwdInterceptor))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewGiveTakeServiceClient(conn)
	req := &pb.InsertLoginPasswordReq{Type: data.TechData.Type, Title: data.TechData.Title, Login: data.Login,
		Password: data.Password, Tag: data.TechData.Tag, Comment: data.TechData.Comment}

	resp, err := client.InsLogPwd(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Title, err
}

// GetLogPwd - getting a login/password pair from server by record id.
// Accepts the key from LogPwdData map in Storage struct.
// Returns LogPwdModel type.
func (c *ClientGRPC) GetLogPwd(keyStorage int) (models.LogPwdModel, error) {
	result := models.LogPwdModel{}
	idRecord := c.Config.Storage.GetIdRecord(keyStorage, datatypes.LoginPasswordDataType)
	ctx := c.getContext()
	conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds), grpc.WithUnaryInterceptor(c.GetLogPwdInterceptor))
	if err != nil {
		return result, err
	}
	defer conn.Close()

	client := pb.NewGiveTakeServiceClient(conn)
	req := &pb.GetItemReq{Id: idRecord}
	resp, err := client.GetLogPwd(ctx, req)
	if err != nil {
		return result, err
	}
	result.Login = resp.Login
	result.Password = resp.Password
	result.TechData.Type = datatypes.LoginPasswordDataType
	result.TechData.Title = resp.Title
	result.TechData.Tag = resp.Tag
	result.TechData.Comment = resp.Comment
	return result, nil
}

// SendCard - send card data to the server.
func (c *ClientGRPC) SendCard(data models.CardModel) (string, error) {
	ctx := c.getContext()
	conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds), grpc.WithUnaryInterceptor(c.SendCardInterceptor))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewGiveTakeServiceClient(conn)
	req := &pb.InsertCardReq{Type: data.TechData.Type, Title: data.TechData.Title, Num: data.Num,
		Date: data.Date, Cvc: data.CVC, FirstName: data.FirstName, LastName: data.LastName,
		Tag: data.TechData.Tag, Comment: data.TechData.Comment}

	resp, err := client.InsCard(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Title, err
}

// GetCard - getting a card data from the server by record id.
// Accepts the key from CardData map in Storage struct.
// Returns CardModel type.
func (c *ClientGRPC) GetCard(keyStorage int) (models.CardModel, error) {
	result := models.CardModel{}
	idRecord := c.Config.Storage.GetIdRecord(keyStorage, datatypes.CardDataType)
	ctx := c.getContext()
	conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds), grpc.WithUnaryInterceptor(c.GetCardInterceptor))
	if err != nil {
		return result, err
	}
	defer conn.Close()

	client := pb.NewGiveTakeServiceClient(conn)
	req := &pb.GetItemReq{Id: idRecord}
	resp, err := client.GetCard(ctx, req)
	if err != nil {
		return result, err
	}
	result.Num = resp.Num
	result.Date = resp.Date
	result.CVC = resp.Cvc
	result.FirstName = resp.FirstName
	result.LastName = resp.LastName
	result.TechData.Type = datatypes.CardDataType
	result.TechData.Title = resp.Title
	result.TechData.Tag = resp.Tag
	result.TechData.Comment = resp.Comment
	return result, nil
}

// SendText - send text data to the server.
func (c *ClientGRPC) SendText(data models.TextDataModel) (string, error) {
	ctx := c.getContext()
	conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds), grpc.WithChainUnaryInterceptor(c.SendTextInterceptor))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewGiveTakeServiceClient(conn)
	req := &pb.InsertTextReq{Type: data.TechData.Type, Title: data.TechData.Title,
		Data: data.Data, Tag: data.TechData.Tag, Comment: data.TechData.Comment}

	resp, err := client.InsText(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Title, err
}

// GetText- getting a text data from server by record id.
// Accepts the key from TextData map in Storage struct.
// Returns TextModel type.
func (c *ClientGRPC) GetText(keyStorage int) (models.TextDataModel, error) {
	result := models.TextDataModel{}
	idRecord := c.Config.Storage.GetIdRecord(keyStorage, datatypes.TextDataType)
	ctx := c.getContext()
	conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds), grpc.WithUnaryInterceptor(c.GetTextInterceptor))
	if err != nil {
		return result, err
	}
	defer conn.Close()

	client := pb.NewGiveTakeServiceClient(conn)
	req := &pb.GetItemReq{Id: idRecord}
	resp, err := client.GetText(ctx, req)
	if err != nil {
		return result, err
	}
	result.Data = resp.Data
	result.TechData.Type = datatypes.TextDataType
	result.TechData.Title = resp.Title
	result.TechData.Tag = resp.Tag
	result.TechData.Comment = resp.Comment
	return result, nil
}

// SendBinary - send binary data to the server.
func (c *ClientGRPC) SendBinary(data models.BinaryDataModel) (string, error) {
	ctx := c.getContext()
	conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds), grpc.WithUnaryInterceptor(c.SendBinaryInterceptor))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewGiveTakeServiceClient(conn)
	req := &pb.InsertBinaryReq{Type: data.TechData.Type, Title: data.TechData.Title,
		Data: data.Data, Tag: data.TechData.Tag, Comment: data.TechData.Comment}

	resp, err := client.InsBinary(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Title, err
}

// GetBinary - getting a binary data from server by record id.
// Accepts the key from BinaryData map in Storage struct.
// Returns BinaryModel type.
func (c *ClientGRPC) GetBinary(keyStorage int) (models.BinaryDataModel, error) {
	result := models.BinaryDataModel{}
	idRecord := c.Config.Storage.GetIdRecord(keyStorage, datatypes.BinaryDataType)
	ctx := c.getContext()
	conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds), grpc.WithUnaryInterceptor(c.GetBinaryInterceptor))
	if err != nil {
		return result, err
	}
	defer conn.Close()

	client := pb.NewGiveTakeServiceClient(conn)
	req := &pb.GetItemReq{Id: idRecord}
	resp, err := client.GetBinary(ctx, req)
	if err != nil {
		return result, err
	}
	result.Data = resp.Data
	result.TechData.Type = datatypes.TextDataType
	result.TechData.Title = resp.Title
	result.TechData.Tag = resp.Tag
	result.TechData.Comment = resp.Comment
	return result, nil
}
