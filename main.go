package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	config "github.com/kcartlidge/simples-config"
)

// Settings ... The loaded ini file contents.
type Settings struct {
	mtx               sync.Mutex
	Port, PollSeconds int
	Endpoints         map[int]config.Entry
}

var (
	settings = Settings{}
	beat     <-chan time.Time
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
	settings.Port = c.GetNumber("SETTINGS", "PORT", 8000)
	settings.PollSeconds = c.GetNumber("SETTINGS", "POLL-SECONDS", 60)
	settings.Endpoints = c.GetSection("ENDPOINTS")
	fmt.Println("Seconds   :", settings.PollSeconds)
	fmt.Println("Endpoints :", len(settings.Endpoints))

	// Process at the expected frequency.
	go func() {
		log.Println("Checking")
		performChecks(settings.Endpoints)

		beat = time.Tick(time.Duration(settings.PollSeconds) * time.Second)
		for _ = range beat {
			log.Println("Checking")
			performChecks(settings.Endpoints)
		}
	}()

	// Start the API server going.
	go func() {
		serve(settings.Port)
	}()

	// Watch for ini file changes and reload.
	go watch("ssd.ini", reload)

	// Wait for enter/return.
	go func() {
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		done <- true
	}()

	fmt.Println()
	fmt.Println("Running on", settings.Port)
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
