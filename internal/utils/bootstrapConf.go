// Package utils bootstraps the api config
package utils

import (
	"log"
	"os"
	"strings"
)

var opts = make(map[string]string)

// GetOpts is called from init()
func GetOpts() map[string]string {
	if len(opts) > 0 {
		return opts
	}
	for _, e := range os.Environ() {
		part := strings.Split(e, "=")
		opts[part[0]] = part[1]
	}
	if opts["HTTP_PORT"] == "" {
		opts["HTTP_PORT"] = "9001"
	}
	if opts["CLAMAV_REST_URL"] == "" {
		opts["CLAMAV_REST_URL"] = "http://localhost:9000"
	}
	if opts["MAX_SIGNATURE_AGE_HOURS"] == "" {
		opts["MAX_SIGNATURE_AGE_HOURS"] = "26"
		log.Printf("MAX_SIGNATURE_AGE_HOURS not set, falling back to %v", opts["MAX_SIGNATURE_AGE_HOURS"])
	}
	if opts["HTTPS_PORT"] != "" {

		if opts["SSL_CERT"] == "" {
			opts["SSL_CERT"] = "/etc/ssl/clamav-rest/server.crt"
		} else {
			log.Printf("Using ssl cert: %v", opts["SSL_CERT"])
		}
		if opts["SSL_KEY"] == "" {
			opts["SSL_KEY"] = "/etc/ssl/clamav-rest/server.key"
		} else {
			log.Printf("Using ssl key: %v", opts["SSL_KEY"])
		}

	}
	return opts
}
