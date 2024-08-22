package main

import (
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	_ "github.com/yukia3e/cloud-functions-monorepo" //nolint:unused-parameter
	"github.com/yukia3e/cloud-functions-monorepo/internal/config"
)

func main() {
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := config.GetPort(); envPort != "" {
		port = envPort
	}

	// By default, listen on all interfaces. If testing locally, run with
	// LOCAL_ONLY=true to avoid triggering firewall warnings and
	// exposing the server outside of your own machine.
	hostname := ""
	if config.IsLocal() {
		hostname = "0.0.0.0"
	}
	if err := funcframework.StartHostPort(hostname, port); err != nil {
		err = fmt.Errorf("failed to start host: %w", err)
		log.Fatalf("funcframework.StartHostPort: %v", err)
	}
}
