package gate

import (
	"context"
	"ichat/internal/format"
	"log"
)

var Map = map[string]func(ctx context.Context, r *format.R){}

func Router(ctx context.Context, r *format.R) {
	if fuc, ok := Map[r.Route]; ok {
		fuc(ctx, r)
	} else {
		log.Print("NOT ROUTER")
	}
}
