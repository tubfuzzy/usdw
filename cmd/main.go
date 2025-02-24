package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"usdw/pkg/server"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// @title USDW API
// @version 1.0
// @description API documentation for USDW application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:3000
// @BasePath /api/v1
func main() {
	serv, err := server.New()
	if err != nil {
		panic(err)
	}

	//if err := serv.App().Listen(serv.Config().Server.Port); err != nil {
	//	serv.Logger().Fatalf("%s", err)
	//}
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	//

	//<-quit

	// Channel for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := serv.App().Listen(serv.Config().Server.Port); err != nil {
			serv.Logger().Errorf("Server error: %s", err)
			quit <- os.Interrupt // Trigger shutdown if Listen fails
		}
	}()

	// Wait for signal
	<-quit

	// err = serv.DB().Close()
	err = serv.Cache().Close()
	err = serv.App().Shutdown()
	if err != nil {
		serv.Logger().Fatalf("%s", err)
	}

	serv.Logger().Info(context.Background(), "Server exited gracefully")

}
