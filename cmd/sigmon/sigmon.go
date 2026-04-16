package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/arizon-dread/clamav-rest-sigmon/api"
	"github.com/arizon-dread/clamav-rest-sigmon/internal/utils"
)

var opts = make(map[string]string)

func init() {
	opts = utils.GetOpts()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health/signature-age", api.SignHandler)

	go serveTLS(mux)
	// Launch server on HTTP
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", opts["SIGMON_HTTP_PORT"]),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Printf("Starting sigmon on port %s", opts["SIGMON_HTTP_PORT"])
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Http server shut down unexpectedly due to error, %v", err)
	}
}

func serveTLS(mux *http.ServeMux) {
	if _, err := os.Stat(opts["SSL_CERT"]); err != nil {
		log.Printf("SSL_CERT not specified, will not run TLS server")
		return
	}
	if _, err := os.Stat(opts["SSL_KEY"]); err != nil {
		log.Printf("SSL_KEY not specified, will not run TLS server")
		return
	}

	tlsSrv := &http.Server{
		Addr:         fmt.Sprintf(":%s", opts["SIGMON_HTTPS_PORT"]),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	if err := tlsSrv.ListenAndServeTLS(opts["SSL_CERT"], opts["SSL_KEY"]); err != nil {
		log.Fatalf("HTTPS server shut down unexpectedly, %v", err)
	}
}
