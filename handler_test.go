package hostname_handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var fn = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hey world!")
	return
}

func TestHostnameHandler(t *testing.T) {
	srv := httptest.NewServer(
		HostnameHeaderHandler(http.HandlerFunc(fn)),
	)
	defer srv.Close()

	//srv.Start()
	// リクエストの送信先はテストサーバのURLへ。
	r, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	resHeader := r.Header.Get(defaultHeaderKey)
	expect, _ := os.Hostname()
	if resHeader != expect {
		t.Errorf("Faile. (response, expect) = (%s, %s)\n", resHeader, expect)
	}
}
