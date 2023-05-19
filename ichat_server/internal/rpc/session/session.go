package session

import "ichat/pkg/protocol/pb"

type session struct{}

var Session = new(session)

func (*session) List(r *pb.SessionListGetRequest) {

}
