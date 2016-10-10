# go-hostname-handler

[![GoDoc](https://godoc.org/github.com/yuokada/go-hostname-handler?status.svg)](https://godoc.org/github.com/yuokada/go-hostname-handler)
[![Build Status](https://travis-ci.org/yuokada/go-hostname-handler.svg?branch=master)](https://travis-ci.org/yuokada/go-hostname-handler)
[![Go Report Card](https://goreportcard.com/badge/github.com/yuokada/go-hostname-handler)](https://goreportcard.com/report/github.com/yuokada/go-hostname-handler)

Add `X-Hostname` Header in a Response.

## Example

```
func HeyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hey world!")
	return
}

func main() {
	http.HandleFunc("/hey", hostname.HostnameHeaderHandler(
		http.HandlerFunc(HeyHandler),
	))
	http.ListenAndServe(":8001", nil)
}
```
