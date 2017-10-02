package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/kcartlidge/simples-config"
)

func main() {
	// Capture ctrl-c etc.
	done := make(chan bool)
	handleShutdown(done)

	// Banner.
	fmt.Println()
	fmt.Println(" __  __  __ ")
	fmt.Println("(__ (__ |  \\")
	fmt.Println(" __) __)|__/")
	fmt.Println()
	fmt.Println("Simple Service Discovery")
	fmt.Println()

	// Read the config.
	c, err := config.CreateConfig("ssd.ini")
	if err != nil {
		log.Fatalln(err)
	}
	pt := c.GetNumber("SETTINGS", "PORT", 8000)
	ps := c.GetNumber("SETTINGS", "POLL-SECONDS", 60)
	fmt.Println("Seconds   :", ps)
	ep := c.GetSection("ENDPOINTS")
	fmt.Println("Endpoints :", len(ep))

	// Process at the expected frequency.
	go func() {
		log.Println("Checking")
		performChecks(ep)

		c := time.Tick(time.Duration(ps) * time.Second)
		for _ = range c {
			log.Println("Checking")
			performChecks(ep)
		}
	}()

	// Start the API server going.
	go func() {
		serve(pt)
	}()

	// Wait for enter/return.
	go func() {
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		done <- true
	}()

	fmt.Println()
	fmt.Println("Running on", pt)
	fmt.Println("Press enter/return to stop")
	fmt.Println()

	// Wait until finished.
	<-done
	fmt.Println("Stopped")
}

func handleShutdown(done chan bool) {
	// Create a channel and listen on it.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine to allow capture of Ctrl-C.
	go func() {
		_ = <-sigs
		fmt.Println()
		fmt.Println("Termination signalled")
		done <- true
	}()
}
