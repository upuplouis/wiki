package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"handlers"
	"net"
	"net/http"
	"protos"
	"strconv"
	"sync"
	"utils"
)

func startGrpcServer(grpcPort int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	protos.RegisterGetwayServiceServer(grpcServer, handlers.NewService())
	return grpcServer.Serve(lis)
}

func startHttpServer(httpPort int, grpcEndpoint string) error {
	ctx := context.Background()
	ctx, cancle := context.WithCancel(ctx)
	defer cancle()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := protos.RegisterGetwayServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		return nil
	}
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodOptions,
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)
	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), handler)
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	grpcPort, _ := strconv.Atoi(utils.GetValueFromIni("port", "grpc_port"))
	httpPort, _ := strconv.Atoi(utils.GetValueFromIni("port", "http_port"))
	go func() {
		defer wg.Done()
		err := startGrpcServer(grpcPort)
		if err != nil {
			logrus.Fatal(err)
		}
	}()
	go func() {
		defer wg.Done()
		err := startHttpServer(httpPort, fmt.Sprintf("localhost:%d", grpcPort))
		if err != nil {
			logrus.Fatal(err)
		}
	}()
	wg.Wait()
}
