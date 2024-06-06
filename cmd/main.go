package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	somniumSys "somnium/internal/api"
	"somnium/internal/module"
	"somnium/libs/jwt"
	mw "somnium/libs/middleware"
	desc "somnium/pkg/api/somnium/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	var (
		ctx              = context.Background()
		svcGrpc          = getenv("GRPC_PORT")
		svcHttp          = getenv("HTTP_PORT")
		postgresLogin    = getenv("POSTGRES_LOGIN")
		postgresPassword = getenv("POSTGRES_PASSWORD")
		postgresHost     = getenv("POSTGRES_HOST")
		postgresDatabase = getenv("POSTGRES_DATABASE")
		postgresPort     = getenv("POSTGRES_PORT", "5432")
		jwtsecret        = getenv("JWT_SECRET")
		postgresURL      = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", postgresLogin, postgresPassword, postgresHost, postgresPort, postgresDatabase)
		jwtcli           = jwt.NewClient(jwtsecret)
	)

	log.Printf("Is it working?")
	pg, err := pgxpool.Connect(ctx, postgresURL)
	if err != nil {
		panic(err)
	}
	module := module.New(module.WithPostgres(pg),
		module.WithJWT(jwtcli),
	)

	lis, err := net.Listen("tcp", svcGrpc)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	mw := mw.NewMiddleware(jwtcli, pg)
	//The next interceptor Validate
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		mw.AuthInterceptor,
	))
	reflection.Register(grpcServer)
	desc.RegisterSomniumServiceServer(grpcServer, somniumSys.NewSomniumSystem(module))
	log.Printf("Somnium System listening at %v port", svcGrpc)

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := desc.RegisterSomniumServiceHandlerFromEndpoint(ctx, grpcMux, svcGrpc, opts); err != nil {
		log.Fatalf("failed to register: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	fs := http.FileServer(http.Dir("./docs/swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))
	withCors := cors.New(cors.Options{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(mux)
	gwServer := &http.Server{
		Addr:    svcHttp,
		Handler: withCors,
	}
	log.Fatalln(gwServer.ListenAndServe())
}

func getenv(env string, fallback ...string) string {
	value := os.Getenv(env)
	if value != "" {
		return value
	}

	if len(fallback) > 0 {
		value = fallback[0]
	}
	return value
}
