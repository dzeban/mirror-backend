package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	delay := flag.Int("delay", 0, "delay for each request in seconds")
	errors := flag.Bool("errors", false, "inject 500 errors in requests")
	port := flag.Int("port", 20000, "port to listen")
	verbose := flag.Bool("v", false, "log every request")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if *errors {
			if rand.Intn(10)%3 == 0 {
				w.WriteHeader(500)
				return
			}
		}
		time.Sleep(time.Second * time.Duration(*delay))
		if *verbose {
			log.Printf("addr=%s host=%s uri=\"%s\"\n",
				r.RemoteAddr, r.Host, r.RequestURI)
		}
		fmt.Fprintln(w, "Hello")
	})

	log.Printf("Listening on port %d, delay is %d, error injecting is %t\n",
		*port, *delay, *errors)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
