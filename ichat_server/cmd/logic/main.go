package main

/*var client = map[string]string{
	"gateway": ichat.GrpcConf.GatewayListenAddr,
	"logic":   ichat.GrpcConf.LogicListenAddr,
}*/

func main() {
	/*server := grpc.NewServer()
	pb.RegisterLogicServer(server, &api.LogicServer{})
	pb.RegisterRelationServer(server, &api.RelationServer{})
	listen, err := net.Listen("tcp", ichat.GlobalConf.LogicListenAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at %v", ichat.GrpcConf.LogicListenAddr)
	go func() {
		if err = server.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	Dial()

	for {
		time.Sleep(1 * time.Second)
	}*/
}
func ListenLogic() {

}

/*
func Dial() {
	for name, addr := range client {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		switch name {
		case "gateway":
			ichat.RpcClient.Gateway = pb.NewGatewayClient(conn)
			break
		case "logic":
			ichat.RpcClient.Logic = pb.NewLogicClient(conn)
			break
		}
	}
}*/
