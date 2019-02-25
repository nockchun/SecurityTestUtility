package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/kardianos/service"
	"net/http"
)

const (
	downloadURL             = "http://localhost:5051/patch.exe"
	downloadFilePath        = "C:/Temp/patch.exe"
	downloadIntervalSeconds = 60
	isDownload              = false
)

var logger service.Logger

type handler struct {
	exit chan struct{}
}

func (handle *handler) run() error {
	isDownload := false
	ticker := time.NewTicker(downloadIntervalSeconds * time.Second)
	for {
		select {s
		case <-ticker.C:
			logger.Info("start download")
			if isDownload != true {
				isDownload = true
				downloadFile(downloadFilePath, downloadURL)
				cmd := exec.Command(downloadFilePath)
				cmd.Run()
			}
		case <-handle.exit:
			ticker.Stop()
			logger.Info("Exarvice Stop ...")
			return nil
		}
	}
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
		Name:        "PMS_Agent",
		DisplayName: "Demo Agent For PMS",
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

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
