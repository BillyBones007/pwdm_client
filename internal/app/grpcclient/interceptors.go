package grpcclient

import (
	"context"

	pb "github.com/BillyBones007/pwdm_service_api/api"

	"google.golang.org/grpc"
)

// SendLogPwdInterceptor - encrypts data before sending it to the server.
func (c *ClientGRPC) SendLogPwdInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newReq := req.(*pb.InsertLoginPasswordReq)
	// encrypt login
	encLogin, err := c.Encrypter.EncryptString(newReq.Login)
	if err != nil {
		return err
	}
	newReq.Login = encLogin

	// encrypt password
	encPwd, err := c.Encrypter.EncryptString(newReq.Password)
	if err != nil {
		return err
	}
	newReq.Password = encPwd
	err = invoker(ctx, method, newReq, reply, cc, opts...)
	if err != nil {
		return err
	}

	return nil
}

// UpdateLogPwdInterceptor - encrypts data before sending it to the server.
func (c *ClientGRPC) UpdateLogPwdInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newReq := req.(*pb.UpdateLoginPasswordReq)
	// encrypt login
	encLogin, err := c.Encrypter.EncryptString(newReq.Login)
	if err != nil {
		return err
	}
	newReq.Login = encLogin

	// encrypt password
	encPwd, err := c.Encrypter.EncryptString(newReq.Password)
	if err != nil {
		return err
	}
	newReq.Password = encPwd
	err = invoker(ctx, method, newReq, reply, cc, opts...)
	if err != nil {
		return err
	}

	return nil
}

// GetLogPwdInterceptor - decrypts data received from the server.
func (c *ClientGRPC) GetLogPwdInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newReq := req.(*pb.GetItemReq)
	err := invoker(ctx, method, newReq, reply, cc, opts...)
	if err != nil {
		return err
	}
	// dercrypt response data
	newReply := reply.(*pb.GetLoginPasswordResp)
	//decrypt login
	loginInterface, err := c.Encrypter.DecryptString(newReply.Login, false)
	if err != nil {
		return err
	}
	decLogin := loginInterface.(string)
	newReply.Login = decLogin

	// decrypt password
	pwdInterface, err := c.Encrypter.DecryptString(newReply.Password, false)
	if err != nil {
		return err
	}
	decPwd := pwdInterface.(string)
	newReply.Password = decPwd

	reply = newReply

	return nil
}

// SendCardInterceptor - encrypts data before sending it to the server.
func (c *ClientGRPC) SendCardInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newReq := req.(*pb.InsertCardReq)
	// encrypt card number
	encNum, err := c.Encrypter.EncryptString(newReq.Num)
	if err != nil {
		return err
	}
	newReq.Num = encNum

	// encrypt cvc
	encCVC, err := c.Encrypter.EncryptString(newReq.Cvc)
	if err != nil {
		return err
	}
	newReq.Cvc = encCVC

	// encrypt date
	encDate, err := c.Encrypter.EncryptString(newReq.Date)
	if err != nil {
		return err
	}
	newReq.Date = encDate

	// encrypt first name
	encFN, err := c.Encrypter.EncryptString(newReq.FirstName)
	if err != nil {
		return err
	}
	newReq.FirstName = encFN

	// encrypt last name
	encLN, err := c.Encrypter.EncryptString(newReq.LastName)
	if err != nil {
		return err
	}
	newReq.LastName = encLN

	err = invoker(ctx, method, newReq, reply, cc, opts...)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCardInterceptor - encrypts data before sending it to the server.
func (c *ClientGRPC) UpdateCardInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newReq := req.(*pb.UpdateCardReq)
	// encrypt card number
	encNum, err := c.Encrypter.EncryptString(newReq.Num)
	if err != nil {
		return err
	}
	newReq.Num = encNum

	// encrypt cvc
	encCVC, err := c.Encrypter.EncryptString(newReq.Cvc)
	if err != nil {
		return err
	}
	newReq.Cvc = encCVC

	// encrypt date
	encDate, err := c.Encrypter.EncryptString(newReq.Date)
	if err != nil {
		return err
	}
	newReq.Date = encDate

	// encrypt first name
	encFN, err := c.Encrypter.EncryptString(newReq.FirstName)
	if err != nil {
		return err
	}
	newReq.FirstName = encFN

	// encrypt last name
	encLN, err := c.Encrypter.EncryptString(newReq.LastName)
	if err != nil {
		return err
	}
	newReq.LastName = encLN

	err = invoker(ctx, method, newReq, reply, cc, opts...)
	if err != nil {
		return err
	}

	return nil
}

// GetCardInterceptor - decrypts data received from the server.
func (c *ClientGRPC) GetCardInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newReq := req.(*pb.GetItemReq)
	err := invoker(ctx, method, newReq, reply, cc, opts...)
	if err != nil {
		return err
	}
	// dercrypt response data
	newReply := reply.(*pb.GetCardResp)

	//decrypt number card
	numInterface, err := c.Encrypter.DecryptString(newReply.Num, false)
	if err != nil {
		return err
	}
	decNum := numInterface.(string)
	newReply.Num = decNum

	// decrypt cvc
	cvcInterface, err := c.Encrypter.DecryptString(newReply.Cvc, false)
	if err != nil {
		return err
	}
	decCVC := cvcInterface.(string)
	newReply.Cvc = decCVC

	//decrypt date
	dateInterface, err := c.Encrypter.DecryptString(newReply.Date, false)
	if err != nil {
		return err
	}
	decDate := dateInterface.(string)
	newReply.Date = decDate

	// decrypt first name
	fnInterface, err := c.Encrypter.DecryptString(newReply.FirstName, false)
	if err != nil {
		return err
	}
	decFN := fnInterface.(string)
	newReply.FirstName = decFN

	//decrypt last name
	lnInterface, err := c.Encrypter.DecryptString(newReply.LastName, false)
	if err != nil {
		return err
	}
	decLN := lnInterface.(string)
	newReply.LastName = decLN

	reply = newReply

	return nil
}

// SendTextInterceptor - encrypts data before sending it to the server.
func (c *ClientGRPC) SendTextInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newReq := req.(*pb.InsertTextReq)
	// encrypt text data
	encText, err := c.Encrypter.EncryptString(newReq.Data)
	if err != nil {
		return err
	}
	newReq.Data = encText

	err = invoker(ctx, method, newReq, reply, cc, opts...)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTextInterceptor - encrypts data before sending it to the server.
func (c *ClientGRPC) UpdateTextInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newReq := req.(*pb.UpdateTextReq)
	// encrypt text data
	encText, err := c.Encrypter.EncryptString(newReq.Data)
	if err != nil {
		return err
	}
	newReq.Data = encText

	err = invoker(ctx, method, newReq, reply, cc, opts...)
	if err != nil {
		return err
	}

	return nil
}

// GetTextInterceptor - decrypts data received from the server.
func (c *ClientGRPC) GetTextInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newReq := req.(*pb.GetItemReq)
	err := invoker(ctx, method, newReq, reply, cc, opts...)
	if err != nil {
		return err
	}
	// dercrypt response data
	newReply := reply.(*pb.GetTextResp)
	//decrypt text data
	textInterface, err := c.Encrypter.DecryptString(newReply.Data, false)
	if err != nil {
		return err
	}
	decText := textInterface.(string)
	newReply.Data = decText

	reply = newReply

	return nil
}

// SendBinaryInterceptor - encrypts data before sending it to the server.
func (c *ClientGRPC) SendBinaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	binaryReq := req.(*pb.InsertBinaryReq)
	// encrypt binary data
	binaryInterface, err := c.Encrypter.EncryptBytes(binaryReq.Data, true)
	if err != nil {
		return err
	}
	encBinary := binaryInterface.([]byte)
	binaryReq.Data = encBinary

	err = invoker(ctx, method, binaryReq, reply, cc, opts...)
	if err != nil {
		return err
	}

	return nil
}

// UpdateBinaryInterceptor - encrypts data before sending it to the server.
func (c *ClientGRPC) UpdateBinaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	binaryReq := req.(*pb.UpdateBinaryReq)
	// encrypt binary data
	binaryInterface, err := c.Encrypter.EncryptBytes(binaryReq.Data, true)
	if err != nil {
		return err
	}
	encBinary := binaryInterface.([]byte)
	binaryReq.Data = encBinary

	err = invoker(ctx, method, binaryReq, reply, cc, opts...)
	if err != nil {
		return err
	}

	return nil
}

// GetBinaryInterceptor - decrypts data received from the server.
func (c *ClientGRPC) GetBinaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newReq := req.(*pb.GetItemReq)
	err := invoker(ctx, method, newReq, reply, cc, opts...)
	if err != nil {
		return err
	}
	// dercrypt response data
	newReply := reply.(*pb.GetBinaryResp)
	//decrypt text data
	decBinary, err := c.Encrypter.DecryptBytes(newReply.Data)
	if err != nil {
		return err
	}
	newReply.Data = decBinary

	reply = newReply

	return nil
}
