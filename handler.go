package hostname

import (
	"log"
	"net/http"
	"os"
)

const (
	defaultHandlerName = "HostnameHeaderHandler"
	defaultHeaderKey   = "X-Hostname"
)

// HostnameHeaderHandler is that add X-Hostname header in a Response.
func HostnameHeaderHandler(f http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			log.Printf("ERROR: %s : %s\n", defaultHandlerName, err)
		}
		w.Header().Set(defaultHeaderKey, hostname)
		f.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
