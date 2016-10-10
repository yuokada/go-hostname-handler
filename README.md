# go-hostname-handler

Add `X-Hostname` Header in a Response.

## Example

```
func HeyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hey world!")
	return
}

func main() {
	http.HandleFunc("/hey", hostname.HostnameHandler(
		http.HandlerFunc(HeyHandler),
	))
	http.ListenAndServe(":8001", nil)
}
```
