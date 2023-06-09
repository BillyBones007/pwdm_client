package grpcclient

import (
	pb "github.com/BillyBones007/pwdm_service_api/api"
	"google.golang.org/grpc"
)

// DeleteRecord - delete record from server.
// Accepts the key from some map in Storage structure, and data type.
// Returns error or nil.
func (c *ClientGRPC) DeleteRecord(keyStorage int, dataType int32) error {
	idRecord := c.Config.Storage.GetIdRecord(keyStorage, dataType)
	ctx := c.getContext()
	conn, err := grpc.Dial(c.Config.ServerAddr, grpc.WithTransportCredentials(c.creds))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewDeleteServiceClient(conn)
	req := &pb.DeleteItemReq{Id: idRecord, Type: dataType}
	_, err = client.DelItem(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
