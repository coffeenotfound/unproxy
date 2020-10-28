package main

import (
	"net/http"
	"log"
	"time"
	"github.com/elazarl/goproxy"
)

func main() {
	// Create and config proxy
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false
	
	proxy.OnRequest().DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		timestamp := time.Now()
		log.Printf("[%v] %v\n", timestamp, r.Host)
		return r, nil
	})
	
	// Start proxy
	listenAddr := ":3128"
	log.Printf("Starting unproxy on %v\n", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, proxy))
}
