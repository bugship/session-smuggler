// Package main is the entry point for the session-smuggler REST API.
// It boots the Fiber app via app.SetupAndRunApp and exits on fatal startup errors.
package main

import (
	"log"

	"github.com/bugship/session-smuggler/app"
)

func main() {
	// SetupAndRunApp loads env, connects to Postgres, registers routes, and listens.
	if err := app.SetupAndRunApp(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}
}
