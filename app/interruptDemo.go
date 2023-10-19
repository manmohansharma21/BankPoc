package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func (s *APIServer) handleInterrupts(w http.ResponseWriter, r *http.Request) {
	// Create a context with a timeout for this request.
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Fprintf(w, "Request canceled or timed out")
	case <-time.After(1 * time.Second):
		fmt.Fprintf(w, "Request processed successfully")
	}
}

func releasePort(port int) {
	// Use the 'lsof' command to find the process using the specified port.
	cmd := exec.Command("lsof", "-t", "-i", fmt.Sprintf(":%d", port))
	output, err := cmd.Output()

	if err == nil {
		// If a process is found, try to kill it.
		pidStr := string(output)
		pid := os.Getpid() // Get the PID of the current process
		for _, p := range strings.Fields(pidStr) {
			processPID, _ := strconv.Atoi(p)
			if processPID != pid {
				// Don't kill the current process.
				if err := syscall.Kill(processPID, syscall.SIGTERM); err == nil {
					fmt.Printf("Killed process with PID %s\n", p)
				}
			}
		}
	}
}

func (s *APIServer) interruptDemo() {
	portStr := s.addr //Using the port

	// Extract the numeric part of the port string (remove the leading colon)
	portStr = strings.TrimLeft(portStr, ":")

	// Convert the port string to an integer
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("Invalid port number: %v\n", err)
		return
	}

	// Attempt to release the port before starting the HTTP server.
	releasePort(port)

	// Create a context that listens for the interrupt signal (Ctrl+C).
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// REGISTER the handler function for the "/myroute" path.
	http.HandleFunc("/myroute", s.handleInterrupts)

	// Start your HTTP server.
	server := &http.Server{Addr: s.addr}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Handle the server startup error here.
			log.Fatalf("Server startup error: %v", err)
		}
	}()

	// Wait for the interrupt signal or other termination conditions.
	<-ctx.Done()

	// Shut down the server gracefully.
	if err := server.Shutdown(context.Background()); err != nil {
		// Handle server shutdown error here.
		log.Fatalf("Server shutdown error: %v", err)
	}
}
