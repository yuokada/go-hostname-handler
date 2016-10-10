package hostname_handler

import (
	"log"
	"net/http"
	"os"
)

const (
	DefaultHandlerName = "HostnameHandler"
	DefaultHeaderKey   = "X-Hostname"
)

func HostnameHandler(f http.Handler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			log.Printf("ERROR: %s : %s\n", DefaultHandlerName, err)
		}
		w.Header().Set(DefaultHeaderKey, hostname)
		f.ServeHTTP(w, r)
	}
	return fn
}
