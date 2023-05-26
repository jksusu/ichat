package chat

import (
	"context"
	"ichat/internal/format"
)

type talk struct{}

var Talk = new(talk)

func ToUser(ctx context.Context, r format.R) {

}

func ToGroup(ctx context.Context, r format.R) {

}

func ToRoom(ctx context.Context, r format.R) {

}
