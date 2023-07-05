package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// use localhost if not being told otherwise
	hostFlag := flag.String("host", "localhost", "a string")
	flag.Parse()
	go func() {
		n := 1
		for {
			// make only 3 requests
			if n == 4 {
				continue
			}
			// make url to make request to
			requestURL := fmt.Sprintf("http://%s:8080/rows?n=%d", *hostFlag, n)
			// make reqest, check for errors
			res, err := http.Get(requestURL)
			if err != nil {
				fmt.Printf("client: error making http request: %s\n", err)
				os.Exit(1)
			}
			// get code of successful request
			log.Printf("client: status code: %d\n", res.StatusCode)
			// if it's 200 get the rows of names
			if res.StatusCode == 200 {
				n++
				resBody, err := io.ReadAll(res.Body)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("client: body: %s", resBody)
			}
		}

	}()

	log.Print("client: get n rows client started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Print("client: get n rows client stopped")
}
