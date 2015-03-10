package main

import (
	"github.com/ridecharge/curbformation-version-service/version"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	//async handling stuff for when app gets shutdown
	wg := new(sync.WaitGroup)
	sigc := make(chan os.Signal)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	//cleanup on shutdown
	go func() {
		select {
		case <-sigc:
			log.Println("waiting for async threads to complete before shutdown")
			wg.Wait()
			os.Exit(0)
		}
	}()

	http.ListenAndServe(":8080", nil)
}
