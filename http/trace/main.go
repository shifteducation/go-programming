package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Printf("ConnectDone: network: %s, addr: %s, err: %s\n", network, addr, err)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	//_, err := http.DefaultTransport.RoundTrip(req)

	client := http.Client{}
	_, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
