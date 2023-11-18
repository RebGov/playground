/*
Copyright © 2023 Rebecca [Becci] Govert <becci.govert@gmail.com>

Server Package handles starting and stopping the server as well as returning the response to user
*/
package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	srv                *http.Server
	ErrServiceCanceled = errors.New("failed to cancel server")
)

func StartServer() {
	createServer()
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func StopServer(ctx context.Context) (done chan struct{}) {
	done = make(chan struct{})
	go func() {
		defer close(done)
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("server shut down failed with error [%s]\n", err)
		}
	}()
	return
}

func createServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", writeJoke)
	srv = &http.Server{
		Addr:    ":5002",
		Handler: mux,
	}
	log.Printf("Service is starting on `%s`", srv.Addr)
}

// writeJoke returns the joke or error response if failure
func writeJoke(w http.ResponseWriter, r *http.Request) {
	resp, err := createResponse()
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	}
	_, err = w.Write(resp)
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	}
}

// WaitShutdown waits until is going to die
func WaitShutdown() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-sigc
	log.Printf("signal received [%v] canceling everything\n", s)
}

func Shutdown(cancel context.CancelFunc) {
	if cancel != nil {
		cancel() // <-- Cancelling the context
	}
	ctx, cancelTimeout := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelTimeout()
	doneHTTP := StopServer(ctx)
	err := waitUntilIsDoneOrCanceled(ctx, doneHTTP)
	if err != nil {
		log.Printf("service has stopped by timeout %s\n", err)
	}
	log.Println("Good bye thank you for reviewing")
}

// waitUntilIsDoneOrCanceled it waits until all the dones channels are closed or the context is canceled
func waitUntilIsDoneOrCanceled(ctx context.Context, dones ...chan struct{}) (err error) {
	done := make(chan struct{})
	go func() {
		for _, d := range dones {
			<-d
		}
		close(done)
	}()
	select {
	case <-done:
		log.Println("server is done")
	case <-ctx.Done():
		err = ErrServiceCanceled
		log.Println("server is canceled")
	}
	return
}
