# Changes

Code differences compared to source project demokratos.

## cmd/demo1kratos/wire_gen.go (+4 -2)

```diff
@@ -13,6 +13,7 @@
 	"github.com/orzkratos/demokratos/demo1kratos/internal/data"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/server"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/service"
+	"github.com/orzkratos/pingkratos/serverpingkratos"
 )
 
 // Injectors from wire.go:
@@ -26,8 +27,9 @@
 	greeterRepo := data.NewGreeterRepo(dataData, logger)
 	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
 	greeterService := service.NewGreeterService(greeterUsecase)
-	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
-	httpServer := server.NewHTTPServer(confServer, greeterService, logger)
+	pingService := serverpingkratos.NewPingService(logger)
+	grpcServer := server.NewGRPCServer(confServer, greeterService, pingService, logger)
+	httpServer := server.NewHTTPServer(confServer, greeterService, pingService, logger)
 	app := newApp(logger, grpcServer, httpServer)
 	return app, func() {
 		cleanup()
```

## internal/server/grpc.go (+9 -1)

```diff
@@ -7,10 +7,17 @@
 	v1 "github.com/orzkratos/demokratos/demo1kratos/api/helloworld/v1"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/service"
+	"github.com/orzkratos/pingkratos/clientpingkratos"
+	"github.com/orzkratos/pingkratos/serverpingkratos"
 )
 
 // NewGRPCServer new a gRPC server.
-func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *grpc.Server {
+func NewGRPCServer(
+	c *conf.Server,
+	greeter *service.GreeterService,
+	pingService *serverpingkratos.PingService,
+	logger log.Logger,
+) *grpc.Server {
 	var opts = []grpc.ServerOption{
 		grpc.Middleware(
 			recovery.Recovery(),
@@ -27,5 +34,6 @@
 	}
 	srv := grpc.NewServer(opts...)
 	v1.RegisterGreeterServer(srv, greeter)
+	clientpingkratos.RegisterPingServer(srv, pingService)
 	return srv
 }
```

## internal/server/http.go (+9 -1)

```diff
@@ -7,10 +7,17 @@
 	v1 "github.com/orzkratos/demokratos/demo1kratos/api/helloworld/v1"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/service"
+	"github.com/orzkratos/pingkratos/clientpingkratos"
+	"github.com/orzkratos/pingkratos/serverpingkratos"
 )
 
 // NewHTTPServer new an HTTP server.
-func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
+func NewHTTPServer(
+	c *conf.Server,
+	greeter *service.GreeterService,
+	pingService *serverpingkratos.PingService,
+	logger log.Logger,
+) *http.Server {
 	var opts = []http.ServerOption{
 		http.Middleware(
 			recovery.Recovery(),
@@ -27,5 +34,6 @@
 	}
 	srv := http.NewServer(opts...)
 	v1.RegisterGreeterHTTPServer(srv, greeter)
+	clientpingkratos.RegisterPingHTTPServer(srv, pingService)
 	return srv
 }
```

## internal/service/service.go (+8 -2)

```diff
@@ -1,6 +1,12 @@
 package service
 
-import "github.com/google/wire"
+import (
+	"github.com/google/wire"
+	"github.com/orzkratos/pingkratos/serverpingkratos"
+)
 
 // ProviderSet is service providers.
-var ProviderSet = wire.NewSet(NewGreeterService)
+var ProviderSet = wire.NewSet(
+	NewGreeterService,
+	serverpingkratos.NewPingService,
+)
```

