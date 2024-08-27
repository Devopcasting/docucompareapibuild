package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-ini/ini"
	"golang.org/x/sys/windows/svc"
)

// Service name
const serviceName = "DocCompareService"

// Config structure
type Config struct {
	PythonPath string
	MainPyPath string
	WorkingDir string
}

// DocCompareService implements svc.Handler
type DocCompareService struct {
	config Config
	status svc.Status
	exit   chan struct{}
}

// Start method for the service
func (s *DocCompareService) Start(_ []string) error {
	s.exit = make(chan struct{})
	go s.run() // Run the FastAPI app in a goroutine
	return nil
}

// Run method to start the Python FastAPI application
func (s *DocCompareService) run() {
	// Change working directory
	if err := os.Chdir(s.config.WorkingDir); err != nil {
		fmt.Println("Error changing working directory:", err)
		return
	}

	// Run the Python FastAPI application
	cmd := exec.Command(s.config.PythonPath, s.config.MainPyPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting FastAPI application:", err)
		return
	}

	// Wait for the command to complete
	if err := cmd.Wait(); err != nil {
		fmt.Println("FastAPI application exited with error:", err)
	}
}

// Stop method for the service
func (s *DocCompareService) Stop() error {
	close(s.exit) // Signal to stop the service
	return nil
}

// Execute method to comply with svc.Handler interface
func (s *DocCompareService) Execute(args []string, req <-chan svc.ChangeRequest, ack chan<- svc.Status) (bool, uint32) {
	ack <- s.status // Send the current status
	for {
		select {
		case <-req:
			ack <- s.status // Acknowledge any requests
		case <-s.exit:
			return false, 0 // Exit the service loop
		}
	}
}

// Main function
func main() {
	// Load the config
	cfg, err := ini.Load("project.ini")
	if err != nil {
		fmt.Println("Error loading ini file:", err)
		return
	}

	section := cfg.Section("doc_compare_service")
	config := Config{
		PythonPath: section.Key("python_path").String(),
		MainPyPath: section.Key("main_py_path").String(),
		WorkingDir: section.Key("working_dir").String(),
	}

	// Create the service
	isWindowsService, err := svc.IsWindowsService()
	if err != nil {
		fmt.Println("Error checking if running as a service:", err)
		return
	}

	if isWindowsService {
		// Running as a Windows service
		// You don't need to create a debug.ConsoleLog service instance here
		err = svc.Run(serviceName, &DocCompareService{config: config})
		if err != nil {
			fmt.Println("Error running service:", err)
		}
	} else {
		// Running in debug mode
		fmt.Println("Running in debug mode...")
		svc.Run(serviceName, &DocCompareService{config: config})
	}
}
