package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kardianos/service"
	"net/http"
)

const (
	documentRoot = "C:/Temp/distribute"
	port         = ":5051"
)

var logger service.Logger

type handler struct {
	exit chan struct{}
}

func (handle *handler) run() error {
	http.Handle("/", http.FileServer(http.Dir(documentRoot)))
	http.ListenAndServe(port, nil)

	return nil
}

func (handle *handler) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	handle.exit = make(chan struct{})

	go handle.run()
	return nil
}

func (handle *handler) Stop(s service.Service) error {
	close(handle.exit)
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "PMS_Server",
		DisplayName: "Demo Server For PMS",
		Description: "This is a simple service for pms.",
	}

	// Create Exarvice service
	program := &handler{}
	s, err := service.New(program, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Setup the logger
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal()
	}

	if len(os.Args) > 1 {

		err = service.Control(s, os.Args[1])
		if err != nil {
			fmt.Printf("Failed (%s) : %s\n", os.Args[1], err)
			return
		}
		fmt.Printf("Succeeded (%s)\n", os.Args[1])
		return
	}

	// run in terminal
	s.Run()
}
