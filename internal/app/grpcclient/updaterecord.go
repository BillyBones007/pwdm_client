package grpcclient

import (
	"github.com/BillyBones007/pwdm_client/internal/customerror"
	"github.com/BillyBones007/pwdm_client/internal/datatypes"
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
	pb "github.com/BillyBones007/pwdm_service_api/api"
	"google.golang.org/grpc"
)

// UpdateRecord - update record on the server.
// Accepts the key from some map in Storage structure, data type, and model.
// Returns title of the updated record and error.
func (c *ClientGRPC) UpdateRecord(keyStorage int, dataType int32, model interface{}) (string, error) {
	idRecord := c.Config.Storage.GetIdRecord(keyStorage, dataType)
	ctx := c.getContext()
	switch dataType {
	case datatypes.LoginPasswordDataType:
		conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds), grpc.WithUnaryInterceptor(c.UpdateLogPwdInterceptor))
		if err != nil {
			return "", err
		}
		defer conn.Close()
		client := pb.NewUpdateServiceClient(conn)

		model := model.(models.LogPwdModel)
		req := &pb.UpdateLoginPasswordReq{Id: idRecord, Title: model.TechData.Title, Login: model.Login,
			Password: model.Password, Tag: model.TechData.Tag, Comment: model.TechData.Comment}
		resp, err := client.UpdateLogPwd(ctx, req)
		if err != nil {
			return "", err
		}
		return resp.Title, nil

	case datatypes.CardDataType:
		conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds), grpc.WithUnaryInterceptor(c.UpdateCardInterceptor))
		if err != nil {
			return "", err
		}
		defer conn.Close()
		client := pb.NewUpdateServiceClient(conn)

		model := model.(models.CardModel)
		req := &pb.UpdateCardReq{Id: idRecord, Title: model.TechData.Title, Num: model.Num, Date: model.Date,
			Cvc: model.CVC, FirstName: model.FirstName, LastName: model.LastName, Tag: model.TechData.Tag, Comment: model.TechData.Comment}
		resp, err := client.UpdateCard(ctx, req)
		if err != nil {
			return "", err
		}
		return resp.Title, nil

	case datatypes.TextDataType:
		conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds), grpc.WithUnaryInterceptor(c.UpdateTextInterceptor))
		if err != nil {
			return "", err
		}
		defer conn.Close()
		client := pb.NewUpdateServiceClient(conn)

		model := model.(models.TextDataModel)
		req := &pb.UpdateTextReq{Id: idRecord, Title: model.TechData.Title, Data: model.Data,
			Tag: model.TechData.Tag, Comment: model.TechData.Comment}
		resp, err := client.UpdateText(ctx, req)
		if err != nil {
			return "", err
		}
		return resp.Title, nil

	case datatypes.BinaryDataType:
		conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds), grpc.WithUnaryInterceptor(c.UpdateBinaryInterceptor))
		if err != nil {
			return "", err
		}
		defer conn.Close()
		client := pb.NewUpdateServiceClient(conn)

		model := model.(models.BinaryDataModel)
		req := &pb.UpdateBinaryReq{Id: idRecord, Title: model.TechData.Title, Data: model.Data,
			Tag: model.TechData.Tag, Comment: model.TechData.Comment}
		resp, err := client.UpdateBinary(ctx, req)
		if err != nil {
			return "", err
		}
		return resp.Title, nil

	default:
		return "", customerror.ErrUnknownDataType
	}
}
