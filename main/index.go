package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)

var port = ":8080"
var store = sessions.NewCookieStore([]byte("SĐJDHKSHDSJHD4D5S44D54SD"))

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func index(w http.ResponseWriter, req *http.Request)  {
	session, _ := store.Get(req, "session-name")

	fmt.Printf("Home pdage accessed by %s\n", session)
	fmt.Fprint(w, "Home page here")
}

func main() {
	http.HandleFunc("/", index)

	fmt.Printf("server listening on port %s\n", port)
	http.ListenAndServe(port, nil)
}
