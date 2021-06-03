package grpc

import (
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/server/config"
)

// StartGRPCWeb starts a gRPC-Web server on the given address.
func StartGRPCWeb(grpcSrv *grpc.Server, config config.Config) (*http.Server, error) {

	wrappedServer := grpcweb.WrapServer(grpcSrv)
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-user-agent, x-grpc-web, grpc-status, grpc-message")
		w.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
		if req.Method == "OPTIONS" {
			return
		}

		wrappedServer.ServeHTTP(w, req)
	})

	grpcWebSrv := &http.Server{
		Addr:    config.GRPCWeb.Address,
		Handler: handler,
	}
	if err := grpcWebSrv.ListenAndServe(); err != nil {
		return nil, err
	}
	return grpcWebSrv, nil
}
