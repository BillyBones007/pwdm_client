package grpcclient

import (
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
	pb "github.com/BillyBones007/pwdm_service_api/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// UpdateInfo - updating data from server.
func (c *ClientGRPC) UpdateInfo() ([]models.InfoModel, error) {
	var listData []models.InfoModel
	ctx := c.getContext()
	conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return listData, err
	}
	defer conn.Close()

	client := pb.NewShowInfoServiceClient(conn)
	req := &pb.Empty{}

	resp, err := client.GetInfo(ctx, req)
	if err != nil {
		return listData, err
	}
	for _, v := range resp.Items {
		item := models.InfoModel{}
		item.Id = v.Id
		item.Type = v.Type
		item.Title = v.Title
		item.Tag = v.Tag
		item.Comment = v.Comment
		listData = append(listData, item)
	}

	return listData, nil
}
