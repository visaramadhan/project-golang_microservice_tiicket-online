package main

import ()

func main() {
	lis, err := net.Listen("top", ":8080")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.CategoryServer(grpcServer, &service.CategoryService)
}
