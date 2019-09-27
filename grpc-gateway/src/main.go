package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"handlers"
	"net"
	"net/http"
	"protos"
	"strconv"
	"strings"
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
	mux := runtime.NewServeMux(runtime.WithProtoErrorHandler(ProtoErrorHandler),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
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
	handler = GetTokenServer(handler)

	fileServer := http.StripPrefix("", http.FileServer(http.Dir("")))
	handler = http.StripPrefix("", handler)
	handler = setFileServer(fileServer, handler)

	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), handler)
}

func setFileServer(fileServer, other http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.RequestURI, "") {
			fileServer.ServeHTTP(w, r)
		}else {
			other.ServeHTTP(w, r)
		}
	})
}

func GetTokenServer(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodOptions {
			r.Header.Set("Pragma", strings.TrimSpace(r.Header.Get("token")))
			//md := metadata.Pairs("token", token)
			//ctx := metadata.NewOutgoingContext(r.Context(), md)
			//r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		}else {
			handler.ServeHTTP(w, r)
		}
	})
}

func ProtoErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	s, ok := status.FromError(err)
	if !ok || s.Code() != codes.OK {
		message := new(interface{})
		w.Header().Del("Trailer")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		bs, _ := json.Marshal(message)
		_, err = w.Write(bs)
	}else {
		runtime.DefaultHTTPProtoErrorHandler(ctx, mux, marshaler, w, r, err)
	}
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
