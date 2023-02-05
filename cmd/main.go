package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"atlassian.carcgl.com/bitbucket/ls/lms/pkg/service"
	"github.com/jessevdk/go-flags"
)

var opt struct {
	LMSgRPCServerAddr string `short:"g" env:"GRPC_ADDR" long:"grpc-addr" default:":9094" description:"lms micro service gRPC server address"`
	PostgresDBAddr    string `env:"DB_ADDR" long:"db-addr" default:"localhost" description:"postgres db address"`
}

func main() {
	_, err := flags.ParseArgs(&opt, os.Args)
	if err != nil {
		log.Fatalf("error parsing flags: %v", err)
	}

	lmsService := service.New(opt.LMSgRPCServerAddr, opt.PostgresDBAddr)

	lmsService.Run()

	// Capture Ctrl-C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// server.Shutdown()

}
