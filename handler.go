package hostname_handler

import (
	"log"
	"net/http"
	"os"
)

const (
	defaultHandlerName = "HostnameHandler"
	defaultHeaderKey = "X-Hostname"
)

// HostnameHandler is that add X-Hostname header in a Response.
func HostnameHandler(f http.Handler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			log.Printf("ERROR: %s : %s\n", defaultHandlerName, err)
		}
		w.Header().Set(defaultHeaderKey, hostname)
		f.ServeHTTP(w, r)
	}
	return fn
}
