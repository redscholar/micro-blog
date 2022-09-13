package handler

import (
	"context"
	"io"
	"time"

	log "go-micro.dev/v4/logger"

	pb "web/proto"
)

type Web struct{}

func (e *Web) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received Web.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Web) ClientStream(ctx context.Context, stream pb.Web_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *Web) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Web_ServerStreamStream) error {
	log.Infof("Received Web.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *Web) BidiStream(ctx context.Context, stream pb.Web_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
