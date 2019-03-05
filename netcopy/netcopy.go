package main

import (
	"fmt"
	"log"

	"github.com/devfacet/gocmd"
)

func main() {
	flags := struct {
		Help          bool `short:"h" long:"help" description:"Display usage" global:"true"`
		Version       bool `short:"v" long:"version" description:"Display version"`
		VersionEx     bool `long:"vv" description:"Display version (extended)"`
		InterfaceList struct {
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		} `command:"ints" description:"Find all information of network interfaces."`
		Server struct {
			Int string `short:"i" long:"interface" required:"true" description:"Interface Name for receiving"`
		} `command:"rx" description:"Start server for receiving a data from client(tx-host)."`
		Client struct {
			File     string `short:"p" long:"target" required:"true" description:"Full-Path of file what you want to send."`
			Fragment int    `short:"f" long:"fragments" required:"true" description:"Frigments size of file."`
			Host     string `short:"t" long:"host" required:"true" description:"Where you want to copy."`
		} `command:"tx" description:"Send file to server(rx-host)"`
	}{}

	// Echo command
	gocmd.HandleFlag("InterfaceList", func(cmd *gocmd.Cmd, args []string) error {
		err := printDevice()
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	})

	// Server Rx commands
	gocmd.HandleFlag("Server", func(cmd *gocmd.Cmd, args []string) error {
		fmt.Println("server start at interface of " + flags.Server.Int)
		err := rx(flags.Server.Int)
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	})

	// Client Tx commands
	gocmd.HandleFlag("Client", func(cmd *gocmd.Cmd, args []string) error {
		fmt.Println("Send a file " + flags.Client.File)
		err := tx(flags.Client.File, flags.Client.Fragment, flags.Client.Host)
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	})

	// Init the app
	gocmd.New(gocmd.Options{
		Name:        "netcopy",
		Version:     "0.1.0",
		Description: "A basic app",
		Flags:       &flags,
		ConfigType:  gocmd.ConfigTypeAuto,
	})
}
