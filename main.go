package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"ditto.co.jp/grpc/pb"
	"ditto.co.jp/grpc/service"
	"ditto.co.jp/grpc/single"
	"github.com/jessevdk/go-flags"
	"google.golang.org/grpc"
)

//options - parameter
type options struct {
	Port int `short:"p" long:"port" description:"service port"`
}

func main() {
	var uid string
	var opts options
	if cmds, err := flags.Parse(&opts); err != nil {
		os.Exit(-1)
	} else {
		if len(cmds) > 0 {
			uid = cmds[0]
		}
	}
	if uid != pb.UUID {
		// fmt.Println("This program can only be launched from the s3service")
		fmt.Println("This program can only be run from within the s3service")
		os.Exit(0)
	}

	guid := fmt.Sprintf("Global\\{%v}", uid)
	//a single instance (windows only)
	sp := single.New(guid)
	err := sp.Open()
	if err != nil {
		os.Exit(0)
	}

	if opts.Port == 0 {
		opts.Port = 9828
	}
	opts.Port = opts.Port + 1
	listenon := fmt.Sprintf("localhost:%v", opts.Port)
	port, err := net.Listen("tcp", listenon)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//new server
	server := grpc.NewServer()
	svc := &service.LocalService{}
	go func() {
		//resigter
		pb.RegisterLocalServiceServer(server, svc)
		fmt.Printf("listening on %v\n", listenon)
		server.Serve(port)
	}()

	//Quit
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("Ctrl+C pressed in Terminal\n")
	server.Stop()

	os.Exit(0)
}
