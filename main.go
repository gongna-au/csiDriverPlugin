package main

import (
	"context"
	"log"
	"net"

	"github.com/csiDriverPlugin/csi"
	"google.golang.org/grpc"
)

type identityServer struct {
	csi.UnimplementedIdentityServer
}

func (*identityServer) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	return &csi.GetPluginInfoResponse{
		Name:          "com.example.myplugin",
		VendorVersion: "0.1.0",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	csi.RegisterIdentityServer(s, &identityServer{})

	log.Println("Starting server on port 50051...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
