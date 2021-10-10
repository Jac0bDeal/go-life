package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jac0bDeal/go-life/internal/life"
	"github.com/Jac0bDeal/go-life/internal/output"
)

func main() {
	halt := make(chan os.Signal, 2)
	signal.Notify(halt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	terminal := output.NewTerminal()
	defer terminal.Close()

	world := life.NewRandomWorld(50, 50, 0.25)
	err := terminal.Print(world.State())
	if err != nil {
		log.Fatalf("failed to write to terminal output: %v", err)
		return
	}

	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-halt:
			return
		case <-ticker.C:
			world.Update()
			err = terminal.Print(world.State())
			if err != nil {
				log.Fatalf("failed to write to terminal output: %v", err)
				return
			}
		}
	}
}
