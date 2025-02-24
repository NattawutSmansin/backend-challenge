package router

import (
	"backend-challenge/module/thirdChallenge/handler"
	initialize "backend-challenge/module/thirdChallenge/initialize"
	protoBuff "backend-challenge/module/thirdChallenge/proto"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// startGRPCConnection เชื่อมต่อกับ gRPC server
func startGRPCConnection(target string, timeout time.Duration) (*grpc.ClientConn, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	conn, err := grpc.DialContext(ctx, target, grpc.WithTransportCredentials(insecure.NewCredentials())) // ใช้ credentials แบบไม่เข้ารหัส
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	return conn, ctx, cancel
}

// setupHTTPGateway สร้าง HTTP Gateway Mux และลงทะเบียน gRPC service
func setupHTTPGateway(ctx context.Context, handler *handler.Handler) *runtime.ServeMux {
	gwmux := runtime.NewServeMux()
	err := protoBuff.RegisterPieFireDireServicesHandlerServer(ctx, gwmux, handler)
	if err != nil {
		log.Fatalf("failed to register handler: %v", err)
	}
	return gwmux
}

// createHTTPServer สร้าง HTTP server ที่ใช้พอร์ต 8080
func createHTTPServer(mux *runtime.ServeMux) *http.Server {
	return &http.Server{
		Addr:    ":8080", // ตั้งพอร์ต 8080 สำหรับ HTTP API
		Handler: mux,     // เชื่อมต่อกับ HTTP Gateway mux
	}
}

// StartGRPCServer ฟังก์ชันหลักในการเริ่มต้น gRPC server และ HTTP server
func StartGRPCServer() {
	target := "localhost:50051"
	conn, ctx, cancel := startGRPCConnection(target, 10*time.Second)
	defer conn.Close()
	defer cancel()

	// สร้าง handler
	handler := initialize.InitializeHandler()

	// สร้าง HTTP Gateway
	gwmux := setupHTTPGateway(ctx, handler)

	// สร้าง HTTP server
	httpServer := createHTTPServer(gwmux)

	// เริ่มต้น HTTP server ที่พอร์ต 8080
	logHttp := fmt.Sprintf("http://localhost%s", httpServer.Addr)
	logEndpoint := fmt.Sprintf("%s/beef/summary", logHttp)

	log.Println("HTTP server running on", logHttp)
	log.Println("Endpoint:", logEndpoint)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
