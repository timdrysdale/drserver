// executable drserver servers digital resources (dr)
package main

import (
	"context"
	"flag"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/timdrysdale/dr/ram"
	"github.com/timdrysdale/dr/restapi"
)

const (
	HOST = "localhost"
)

var listen string

func init() {
	flag.StringVar(&listen, "listen", "8085", "listenting port")
	flag.Parse()
}

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan struct{})

	// configure & start server

	portCheck(listen)

	srv := &http.Server{Addr: ":" + listen}

	ram := ram.New()

	srv.Handler = restapi.New(ram)

	go func() {
		//https://stackoverflow.com/questions/39320025/how-to-stop-http-listenandserve
		// returns ErrServerClosed on graceful close
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.WithField("error", err).Fatal("http.ListenAndServe")
		}
		log.Debug("Exiting http.Server")
	}()

	log.Debug("Started http.Server")

	// wait for end
	go func() {
		<-sigs
		close(done)
	}()

	<-done

	// cleanup goes here

	log.Debug("Starting to close http.Server")

	if err := srv.Shutdown(context.TODO()); err != nil {
		log.WithField("error", err).Fatal("Failure/timeout shutting down the http.Server gracefully")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		log.WithField("error", err).Fatal("Could not gracefully shutdown http.Server")
	}

	log.Debug("Stopped http.Server")

}
