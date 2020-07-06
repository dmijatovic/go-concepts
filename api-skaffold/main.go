package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"./handlers"
)

var certFile string = "/home/dusan/test/go/concepts/grpc/ssl/cert/server.crt"
var keyFile string = "/home/dusan/test/go/concepts/grpc/ssl/cert/server.pem"

func createHTTPSServer(mux http.Handler, host string) {
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         host,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Println("Starting https server on...", host)
	log.Fatal(srv.ListenAndServeTLS(certFile, keyFile))
	// return srv
}

func httpStart() {
	router := handlers.Setup()
	createHTTPSServer(router, ":8080")
}

func main() {
	log.Println("App starting...")
	// start http server
	httpStart()
}
