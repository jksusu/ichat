package grpc

import "ichat/internal/format"

type grpc struct {
	r *format.R
}

func NewRequest(r *format.R) *grpc {
	return &grpc{r: r}
}

func (*grpc) Push() {

}

func (*grpc) Logic() {

}

func (*grpc) Gateway() {

}
