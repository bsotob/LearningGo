package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

var (
	proxyLAN, _ = url.Parse("http://proxynexe.svb.lacaixa.es:8080")
)

func main() {
	var proto string
	flag.StringVar(&proto, "proto", "https", "Protocol HTTP or HTTPS (http by default)")

	log.Println("Setting multiplexor")
	mux := http.NewServeMux()
	mux.HandleFunc("/", logger(forwardRequest))
	log.Println("Mux is setted")

	log.Println("Configuring ForwardProxy server")
	server := &http.Server{
		Addr:         ":8888",
		Handler:      mux,
		TLSConfig:    &tls.Config{InsecureSkipVerify: true},
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	if proto == "http" {
		log.Println("Proxy HTTP running")
		log.Fatal(server.ListenAndServe())
	} else {
		log.Println("Proxy HTTPS runnnig")
		log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
	}
}

func logger(f func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("[ %s ] %s - %s", req.RemoteAddr, req.Method, req.URL.String())
		f(w, req)
	}
}

func copyHeaders(dst, source http.Header) {
	for key, vv := range source {
		for _, value := range vv {
			dst.Add(key, value)
		}
	}
}

func handleTunneling(w http.ResponseWriter, req *http.Request) {
	destConn, err := net.DialTimeout("tcp", req.Host, 5*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusServiceUnavailable)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	go connect(destConn, clientConn)
	go connect(clientConn, destConn)
}

func connect(dest io.WriteCloser, source io.ReadCloser) {
	defer dest.Close()
	defer source.Close()

	io.Copy(dest, source)
}

func forwardRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodConnect {
		log.Println("Starting tunneling")
		handleTunneling(w, req)
	} else {
		var DefaultTransport http.RoundTripper = &http.Transport{
			Proxy: http.ProxyURL(proxyLAN),
			// TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		resp, err := DefaultTransport.RoundTrip(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		copyHeaders(w.Header(), resp.Header)
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)

	}
}
