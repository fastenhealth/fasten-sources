package main

import (
	"github.com/skratchdot/open-golang/open"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Printf("%v", req.URL.Path)
		if strings.HasPrefix(req.URL.Path, "/callback") {
			http.ServeFile(res, req, "html/callback.html")
		} else {
			http.ServeFile(res, req, "html/index.html")
		}
	})
	url := "http://localhost:9999"

	go func() {
		log.Println("You will now be taken to your browser for authentication")
		time.Sleep(1 * time.Second)
		err := open.Run(url)
		if err != nil {
			log.Fatalf("an error occurred opening browser: %v", err)
		}
		time.Sleep(1 * time.Second)
		log.Printf("Authentication URL: %s\n", url)
	}()

	log.Fatal(http.ListenAndServe(":9999", nil))

}
