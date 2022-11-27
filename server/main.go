package main

import (
	"fmt"
	"log"

	"github.com/pocketbase/pocketbase"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(" [🍰] " + string(bytes))
}
func configLogging() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

func main() {
	configLogging()
	log.Println("Welcome!")
	backend := pocketbase.New()
	backend.RootCmd.AddCommand(NewGrpcCommand(backend))
	if err := backend.Start(); err != nil {
		log.Fatal(err)
	}
	log.Println("PocketBase Initialized!")
}
