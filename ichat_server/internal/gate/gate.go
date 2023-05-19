package gate

import (
	"context"
	"ichat/pkg/protocol/pb"
)

type gate struct{}

var Gateway = new(gate)

func (*gate) Login() {

}

func (*gate) Response(ctx context.Context, req *pb.ResponseReq) error {
	return nil
}
