package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tomasen/realip"
)

// Source code:
// https://tachingchen.com/tw/blog/get-real-client-ip-inside-kubernetes-through-http-loadbalancer/

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host := os.Getenv("HOSTNAME")
		reply := fmt.Sprintf("Hostname:\n%s\n\nUser-Agent:\n%v\n\nHeader:\n%v\n\nIP:\n%v", host, r.UserAgent(), r.Header, realip.RealIP(r))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(reply))
	})
	http.ListenAndServe(":80", nil)
}
